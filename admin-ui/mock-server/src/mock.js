const { MockList } = require('apollo-server');
const casual = require('casual');
const { UserActivityOptions, metricsData } = require('./mockSamples');
const { PubSub } = require('apollo-server-express');
const { workflowNames, processNames, emails } = require('./data');

const pubsub = new PubSub();

const generateRuntime = () => ({
  id: parseInt(casual.array_of_digits(8).join('')),
  name: casual.name,
  description: casual.sentences(10),
  status: casual.random_element(['CREATING', 'STARTED', 'ERROR']),
  creationAuthor: () => ({
    id: casual.uuid,
    email: casual.random_element(emails)
  }),
  creationDate: casual.moment.toISOString(),
  versions: () => new MockList([1, 5]),
  publishedVersion: () => {
    if (Math.random() > 0.5) {
      return {
        id: parseInt(casual.array_of_digits(8).join('')),
        name: `v${casual.integer(1, 10)}.${casual.integer(
          (from = 1),
          (to = 10)
        )}.${casual.integer((from = 1), (to = 10))}`,
        description: casual.sentence,
        status: 'STARTED',
        creationDate: casual.moment.toISOString(),
        activationDate: casual.moment.toISOString()
      };
    }
    return null;
  }
});

const generateVersion = () => ({
  id: casual.random_element(['v1', 'v2', 'v3', 'v4', 'v5', 'v6']),
  name: `v${casual.integer((from = 1), (to = 10))}.${casual.integer(
    (from = 1),
    (to = 10)
  )}.${casual.integer((from = 1), (to = 10))}`,
  description: casual.sentence,
  status: casual.random_element([
    'STARTING',
    'STARTED',
    'PUBLISHED',
    'STOPPED'
  ]),
  creationDate: casual.moment.toISOString(),
  publicationDate: casual.moment.toISOString(),
  configurationVariables: () => new MockList([2, 20]),
  workflows: () => new MockList(2),
  configurationCompleted: true
});

function getProcessName() {
  return casual.random_element(processNames);
}
function getWorkflowName() {
  return casual.random_element(processNames);
}

function getLogLevel() {
  const rand = Math.random();
  switch (true) {
    case rand <= 0.02:
      return 'DEBUG';
    case rand <= 0.06:
      return 'ERROR';
    case rand <= 0.12:
      return 'WARN';
    default:
      return 'INFO';
  }
}

function getLogMessage() {
  return casual.sentences(casual.integer(1, 4));
}

let activeNodeLogsInterval = null;

const getPercStr = () => casual.integer((from = 0), (to = 100)).toString();

const FREQUENCY_LOGS = 2000;

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

module.exports = {
  Subscription: () => ({
    nodeLogs: {
      subscribe: () => {
        activeNodeLogsInterval && clearInterval(activeNodeLogsInterval);
        activeNodeLogsInterval = setInterval(() => {
          pubsub.publish('nodeLogs', {
            nodeLogs: {
              id: casual.uuid,
              date: new Date().toUTCString(),
              nodeId: casual.uuid,
              nodeName: getProcessName(),
              workflowId: casual.uuid,
              workflowName: getWorkflowName(),
              message: getLogMessage,
              level: getLogLevel
            }
          });
        }, FREQUENCY_LOGS);
        setTimeout(() => {
          clearInterval(activeNodeLogsInterval);
        }, 5 * FREQUENCY_LOGS);
        return pubsub.asyncIterator('nodeLogs');
      }
    },
    runtimeCreated: {
      subscribe: () => pubsub.asyncIterator('runtimeCreated')
    },
    versionNodeStatus: {
      subscribe: () => pubsub.asyncIterator('versionNodeStatus')
    }
  }),
  Query: () => ({
    me: () => ({
      id: casual.uuid,
      accessLevel: 'ADMIN',
      email: 'admin@intelygenz.com'
    }),
    users: () => new MockList([20, 30]),
    runtimes: () => new MockList([4, 8]),
    alerts: () => new MockList([1, 4]),
    versions: () => new MockList([18, 28]),
    domains: () => new MockList([2, 6]),
    userActivityList: () => new MockList([30, 30]),
    logs: async () => {
      await sleep(2000);
      return {
        cursor: casual.string,
        items: () => new MockList(40)
      };
    }
  }),
  Mutation: () => ({
    createRuntime: () => {
      setTimeout(() => {
        const _runtime = generateRuntime();
        pubsub.publish('runtimeCreated', {
          runtimeCreated: { id: _runtime.id, name: _runtime.name }
        });
      }, 4000);
      return { errors: [], runtime: this.Runtime };
    },
    createVersion: () => ({ errors: [], version: version }),
    updateSettings: () => ({ errors: [], settings: this.Settings })
  }),
  MetricsValues: () => ({
    accuracy: this.MetricsAccuracy,
    missing: getPercStr(),
    newLabels: getPercStr()
  }),
  MetricsAccuracy: () => ({
    total: getPercStr(),
    micro: getPercStr(),
    macro: getPercStr(),
    weighted: getPercStr()
  }),
  MetricsCharts: () => ({
    confusionMatrix: metricsData.DataMatrix(),
    seriesAccuracy: metricsData.DataNumberStr(),
    seriesRecall: metricsData.DataNumberStr(),
    seriesSupport: metricsData.DataNumberStr(),
    successVsFails: metricsData.DataHourNumber()
  }),
  User: () => ({
    id: casual.uuid,
    accessLevel: casual.random_element(['ADMIN', 'VIEWER', 'MANAGER']),
    email: casual.random_element(emails),
    creationDate: new Date().toUTCString(),
    lastActivity: new Date().toUTCString()
  }),
  LogPane: () => ({
    cursor: casual.string,
    items: new MockList([6, 12])
  }),
  NodeLog: () => ({
    id: casual.uuid,
    date: new Date().toUTCString(),
    nodeId: casual.uuid,
    nodeName: getProcessName(),
    workflowId: casual.uuid,
    workflowName: getWorkflowName(),
    message: getLogMessage,
    level: getLogLevel
  }),
  Runtime: generateRuntime,
  Version: generateVersion,
  ConfigurationVariable: () => ({
    key: casual.word.toUpperCase(),
    value: () => {
      if (Math.random() < 0.1) {
        return '';
      }
      return casual.sentence;
    },
    type: casual.random_element(['VARIABLE', 'FILE'])
  }),
  Alert: () => ({
    id: casual.uuid,
    type: 'ERROR',
    message: casual.sentence,
    runtime: this.Runtime
  }),
  UserActivity: () => casual.random_element(UserActivityOptions(this.User)),
  Settings: () => ({
    authAllowedDomains: () => new MockList([2, 6], () => casual.domain),
    sessionLifetimeInDays: () => casual.integer(1, 99)
  }),
  Workflow: () => ({
    name: casual.random_element(workflowNames),
    nodes: () => new MockList([1, 4]),
    edges: () => new MockList([1, 4])
  }),
  Edge: () => ({ id: casual.uuid, fromNode: casual.uuid, toNode: casual.uuid }),
  Node: () => {
    const _id = casual.uuid;
    setTimeout(() => {
      pubsub.publish('versionNodeStatus', {
        versionNodeStatus: {
          date: new Date().toUTCString(),
          nodeId: _id,
          nodeName: getProcessName(),
          status: casual.random_element(['STARTED', 'ERROR']),
          message: 'message'
        }
      });
    }, casual.integer(2000, 10000));
    return { id: _id, name: getProcessName(), status: 'STOPPED' };
  },
  MetricChartData: () => new MockList([5, 10])
};
