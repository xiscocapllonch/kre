import abc
from datetime import datetime
import traceback
import uuid

from grpclib import GRPCError
from protobuf_to_dict import protobuf_to_dict

from kre_nats import KreNatsMessage
from kre_measurements import KreMeasurements


# NOTE: EntrypointKRE will be extended by Entrypoint class auto-generated
class EntrypointKRE:
    def __init__(self, logger, nc, subjects, config):
        self.logger = logger
        self.subjects = subjects
        self.nc = nc
        self.config = config
        self.measurements = KreMeasurements(config)

    @abc.abstractmethod
    async def make_response_object(self, subject, response):
        # To be implemented on autogenerated entrypoint
        pass

    async def process_message(self, stream, subject) -> None:
        start = datetime.utcnow()
        tracking_id = str(uuid.uuid4())

        try:
            raw_msg = await stream.recv_message()
            request_dict = protobuf_to_dict(raw_msg)

            self.logger.info(f'gRPC message received')
            request_msg = KreNatsMessage(data=request_dict, tracking_id=tracking_id)

            nats_subject = self.subjects[subject]
            self.logger.info(f"Starting request/reply on NATS subject: '{nats_subject}'")

            nats_reply = await self.nc.request(nats_subject, request_msg.marshal(),
                                               timeout=self.config.request_timeout)

            self.logger.info(f"creating a response from message reply")
            response_msg = KreNatsMessage(msg=nats_reply)

            if response_msg.error:
                self.logger.error(f"received message: {response_msg}")

            response = self.make_response_object(subject, response_msg)

            await stream.send_message(response)

            self.logger.info(f'gRPC successfully response')

            self.save_elapsed_times(start, tracking_id, subject, response_msg.tracking)

        except Exception as err:
            err_msg = f'Exception on gRPC call : {err}'
            self.logger.error(err_msg)
            traceback.print_exc()

            if isinstance(err, GRPCError):
                raise err
            

    def save_elapsed_times(self, entrypoint_start: datetime, tracking_id: str, subject: str, tracking: list) -> None:
            entrypoint_end = datetime.utcnow()
            elapsed = (entrypoint_end - entrypoint_start).total_seconds() * 1000

            # Save the elapsed time measurement
            fields = {"elapsed_ms": elapsed, "tracking_id": tracking_id}
            tags = {
                "workflow": subject,
                "version": self.config.krt_version,
            }
            self.measurements.save("workflow_elapsed_time", fields, tags)

            # Save the elapsed time for each node
            for idx, t in enumerate(tracking):
                tags = {
                    "workflow": subject,
                    "version": self.config.krt_version,
                    "node": t["node_name"]
                }

                if idx == 0:
                    prev_node_name = "entrypoint"
                    prev_node_end = entrypoint_start
                else:
                    prev_node = tracking[idx-1]
                    prev_node_name = prev_node["node_name"]
                    prev_node_end = datetime.fromisoformat(prev_node["end"])

                node_start = datetime.fromisoformat(t["start"])
                node_end = datetime.fromisoformat(t["end"])
                elapsed = node_end - node_start
                waiting = node_start - prev_node_end
                fields = {
                    "tracking_id": tracking_id,
                    "node_from": prev_node_name,
                    "elapsed_ms": elapsed.total_seconds() * 1000,
                    "waiting_ms": waiting.total_seconds() * 1000
                }
                self.measurements.save("node_elapsed_time", fields, tags)
