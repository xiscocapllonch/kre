import { get } from 'lodash';

import React, { FunctionComponent } from 'react';

import Settings from '../../components/Settings/Settings';

import { loader } from 'graphql.macro';
import { useQuery } from '@apollo/react-hooks';

import { GetUserEmail } from '../../graphql/queries/types/GetUserEmail';

import styles from './Header.module.scss';

const GetUserEmailQuery = loader('../../graphql/queries/getUserEmail.graphql');

type Props = {
  children?: any;
};
const Header: FunctionComponent<Props> = ({ children }) => {
  const { data, error, loading } = useQuery<GetUserEmail>(GetUserEmailQuery);

  if (loading)
    return <div className={styles.splash} data-testid={'splashscreen'} />;

  const username: string = get(data, 'me.email');

  return (
    <header className={styles.container}>
      <img
        className={styles.konstellationsIcon}
        src={'/img/brand/konstellation.png'}
        alt="konstellation logo"
      />
      <div className={styles.customHeaderElements}>{children}</div>
      <Settings label={username} />
    </header>
  );
};

export default Header;
