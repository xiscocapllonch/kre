# Build entrypoint-gen
FROM golang:1.14 as builder

ENV CGO_ENABLED=0

WORKDIR /build
COPY entrypoint-gen/go.* ./
RUN go mod download

COPY entrypoint-gen/* ./

RUN go build -o entrypoint-gen .

# Convert Pipfile.lock to requirements.txt
FROM python:latest as requirement-builder
WORKDIR /app
RUN pip install pipfile-requirements

COPY Pipfile.lock ./
RUN pipfile2req > /app/requirements.txt


# Final image build
FROM ubuntu:18.04

# Fix error installing pipenv:
#   RuntimeError: Click will abort further execution because
#   Python 3 was configured to use ASCII as encoding for the environment.
ENV LC_ALL C.UTF-8
ENV LANG C.UTF-8

ENV PYTHONUNBUFFERED=1

RUN apt update && \
    apt -y install python3.7 python3-pip python3.7-dev && \
    update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.7 1 && \
    update-alternatives --set python3 /usr/bin/python3.7

WORKDIR /app

COPY --from=requirement-builder /app/requirements.txt .
RUN python3 -m pip install -r requirements.txt

COPY src src
COPY entrypoint.sh ./

RUN mkdir src/entrypoint

COPY --from=builder /build/entrypoint-gen .
COPY --from=builder /build/entrypoint.py.tmpl .

ENV PYTHONPATH="/app/src/entrypoint"

CMD ["/app/entrypoint.sh"]
