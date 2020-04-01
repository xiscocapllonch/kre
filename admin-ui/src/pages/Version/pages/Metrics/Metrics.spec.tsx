import React from 'react';
import { shallow } from 'enzyme';
import Metrics from './Metrics';
import { version, runtime } from '../../../../mocks/version';

jest.mock('react-router', () => ({
  useParams: jest.fn(() => ({ runtimeId: '', versionId: '' }))
}));
jest.mock('@apollo/react-hooks', () => ({
  useQuery: jest.fn(() => ({
    data: {},
    loading: false,
    error: '',
    refetch: jest.fn()
  }))
}));

// TODO: test GraphQL data
describe('Metrics', () => {
  let wrapper;

  beforeEach(() => {
    wrapper = shallow(<Metrics runtime={runtime} version={version} />);
  });

  it('matches snapshot', () => {
    expect(wrapper).toMatchSnapshot();
  });
});
