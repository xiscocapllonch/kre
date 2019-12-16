import gql from 'graphql-tag';

export const RUNTIME_CREATED_SUBSCRIPTION = gql`
  subscription runtimeCreated {
    runtimeCreated {
      id
      name
    }
  }
`;
