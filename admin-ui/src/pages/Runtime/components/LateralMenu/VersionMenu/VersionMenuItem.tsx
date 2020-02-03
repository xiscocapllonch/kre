import styles from './VersionMenuItem.module.scss';
import React from 'react';
import { NavLink } from 'react-router-dom';
import IconWarning from '@material-ui/icons/Warning';
import cx from 'classnames';

export type VersionMenuItemProps = {
  label: string;
  to: string;
  exact?: boolean;
  warning?: string;
  Icon: any;
};

function VersionMenuItem({
  label,
  to,
  exact,
  warning,
  Icon
}: VersionMenuItemProps) {
  return (
    <NavLink to={to} activeClassName={styles.active} exact={exact} replace>
      <div className={styles.item}>
        <div className={styles.icon}>
          <Icon className="icon-regular" />
        </div>
        <span>{label}</span>
        {warning && (
          <div title={warning} className={styles.iconWarning}>
            <IconWarning className={cx('icon-regular', styles.warning)} />
          </div>
        )}
      </div>
    </NavLink>
  );
}

export default VersionMenuItem;
