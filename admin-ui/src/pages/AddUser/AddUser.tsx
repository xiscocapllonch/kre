import * as CHECK from '../../components/Form/check';

import {
  CreateUser,
  CreateUserVariables,
  CreateUser_createUser
} from '../../graphql/mutations/types/CreateUser';
import React, { useEffect } from 'react';

import { AccessLevel } from '../../graphql/types/globalTypes';
import Button from '../../components/Button/Button';
import { GetUsers } from '../../graphql/queries/types/GetUsers';
import ROUTE from '../../constants/routes';
import Select from '../../components/Form/Select/Select';
import TextInput from '../../components/Form/TextInput/TextInput';
import { get } from 'lodash';
import { loader } from 'graphql.macro';
import { mutationPayloadHelper } from '../../utils/formUtils';
import styles from './AddUser.module.scss';
import { useForm } from 'react-hook-form';
import { useHistory } from 'react-router';
import { useMutation } from '@apollo/react-hooks';
const GetUsersQuery = loader('../../graphql/queries/getUsers.graphql');

const CreateUserMutation = loader('../../graphql/mutations/createUser.graphql');

function verifyEmail(value: string) {
  return CHECK.getValidationError([
    CHECK.isFieldNotEmpty(value),
    CHECK.isEmailValid(value)
  ]);
}

function verifyAccessLevel(value: string) {
  return CHECK.getValidationError([
    CHECK.isFieldNotEmpty(value),
    CHECK.isFieldNotInList(value, Object.values(AccessLevel))
  ]);
}

type FormData = {
  email: string;
  accessLevel: AccessLevel;
};

function AddUser() {
  const history = useHistory();

  const { handleSubmit, setValue, register, errors, setError, watch } = useForm<
    FormData
  >();
  const [addUser, { loading, error: mutationError }] = useMutation<
    CreateUser,
    CreateUserVariables
  >(CreateUserMutation, {
    onCompleted: onCompleteAddUser,
    update: (cache, { data }) => {
      const newUser = data?.createUser as CreateUser_createUser;
      const cacheResult = cache.readQuery<GetUsers>({
        query: GetUsersQuery
      });

      if (cacheResult !== null) {
        const { users } = cacheResult;
        cache.writeQuery({
          query: GetUsersQuery,
          data: { users: users.concat([newUser]) }
        });
      }
    }
  });

  useEffect(() => {
    register('email', { validate: verifyEmail });
    register('accessLevel', { required: true, validate: verifyAccessLevel });
    setValue('email', '');
    setValue('accessLevel', AccessLevel.VIEWER);
  }, [register, setValue]);

  useEffect(() => {
    if (mutationError) {
      setError('email', '', mutationError.toString());
    }
  }, [mutationError, setError]);

  function onCompleteAddUser(updatedData: CreateUser) {
    history.push(ROUTE.SETTINGS_USERS);
  }

  function onSubmit(formData: FormData) {
    addUser(mutationPayloadHelper(formData));
  }

  return (
    <div className={styles.bg}>
      <div className={styles.grid}>
        <div className={styles.container}>
          <h1>Please introduce a new user</h1>
          <p className={styles.subtitle} />
          <div className={styles.content}>
            <TextInput
              whiteColor
              label="email"
              error={get(errors.email, 'message')}
              onChange={(value: string) => setValue('email', value)}
              onEnterKeyPress={handleSubmit(onSubmit)}
              autoFocus
            />
            <Select
              label="User type"
              showSelectAllOption={false}
              options={Object.values(AccessLevel)}
              onChange={(value: AccessLevel) => setValue('accessLevel', value)}
              error={get(errors.accessLevel, 'message')}
              formSelectedOption={watch('accessLevel')}
              placeholder="Access level"
            />
            <div className={styles.buttons}>
              <Button
                primary
                label="SAVE"
                onClick={handleSubmit(onSubmit)}
                loading={loading}
                className={styles.buttonSave}
              />
              <Button
                label="CANCEL"
                onClick={() => {
                  history.goBack();
                }}
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default AddUser;
