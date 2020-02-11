import React from 'react';
import { renderWithRouter } from './utils/testUtils';
import { cleanup } from '@testing-library/react';
import ROUTE from './constants/routes';
import App, { Routes } from './App';

import { MockedProvider } from '@apollo/react-testing';
import { dashboardMock } from './mocks/runtime';
import { usernameMock } from './mocks/auth';

const mocks = [dashboardMock, usernameMock];

afterEach(cleanup);

it('renders without crashing', () => {
  const {
    element: { container }
  } = renderWithRouter(
    <MockedProvider mocks={mocks} addTypename={false}>
      <App />
    </MockedProvider>
  );
  expect(container).toMatchSnapshot();
});

it('it shows dashboard page on home URL when logged', () => {
  const {
    element: { getByTestId }
  } = renderWithRouter(
    <MockedProvider mocks={mocks} addTypename={false}>
      <Routes />
    </MockedProvider>,
    ROUTE.HOME
  );

  expect(getByTestId('dashboardContainer')).toBeInTheDocument();
});
