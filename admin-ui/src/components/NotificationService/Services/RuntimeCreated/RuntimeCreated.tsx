import { get } from 'lodash';
import React, { useState } from 'react';
import { useHistory, useLocation } from 'react-router';
import Notification from '../../Notification';
import { RUNTIME_CREATED_SUBSCRIPTION } from './RuntimeCreated.graphql';
import { useSubscription } from '@apollo/react-hooks';
import { Runtime } from '../../../../graphql/models';
import ROUTE from '../../../../constants/routes';

const NOTIFICATION_TIMEOUT = 15 * 1000;

type Notification = {
  id: string;
  message: string;
};

function RuntimeCreated() {
  const history = useHistory();
  const location = useLocation();

  const [notifications, setNotifications] = useState<Notification[]>([]);

  useSubscription<Runtime>(RUNTIME_CREATED_SUBSCRIPTION, {
    onSubscriptionData: (msg: any) => {
      const runtime = get(msg, 'subscriptionData.data.runtimeCreated');
      addNotification(runtime);
    }
  });

  function createNotificationObject(runtime: Runtime): Notification {
    return {
      id: runtime.id,
      message: `The RUNTIME "${runtime.name}" has been successfully created!`
    };
  }

  function addNotification(runtime: Runtime) {
    const newNotification = createNotificationObject(runtime);

    // Close notification after NOTIFICATION_TIMEOUT seconds
    setTimeout(() => {
      closeNotification(newNotification.id);
    }, NOTIFICATION_TIMEOUT);

    // Refresh dashboard
    if (location.pathname === ROUTE.HOME) {
      history.push('/other');
      history.replace(ROUTE.HOME);
    }

    setNotifications((notifs: Notification[]) =>
      notifs.concat([newNotification])
    );
  }

  function closeNotification(id: string) {
    setNotifications((notifs: Notification[]) =>
      notifs.filter((notification: Notification) => notification.id !== id)
    );
  }

  const notificationComponents = notifications.map(
    (notification: Notification) => (
      <Notification
        key={`notification_${notification.id}`}
        message={notification.message}
        buttonLabel="GO TO RUNTIME"
        buttonAction={() => {
          const runtimePath = ROUTE.RUNTIME.replace(
            ':runtimeId',
            notification.id
          );
          if (!location.pathname.startsWith(runtimePath)) {
            history.push(runtimePath);
          }

          closeNotification(notification.id);
        }}
        onCloseNotification={() => closeNotification(notification.id)}
      />
    )
  );

  return <div>{notificationComponents}</div>;
}

export default RuntimeCreated;
