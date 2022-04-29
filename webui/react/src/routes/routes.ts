import { RouteConfig } from './types';

const routes: RouteConfig[] = [
  {
    id: 'workspaceDetails',
    needAuth: true,
    path: '/workspaces/:workspaceId',
    title: 'Workspace',
  },
  {
    id: 'workspaceList',
    needAuth: true,
    path: '/workspaces',
    title: 'Workspace',
  },
  {
    id: 'projectDetails',
    needAuth: true,
    path: '/projects/:projectId',
    title: 'Project',
  },
  {
    icon: 'experiment',
    id: 'uncategorized',
    needAuth: true,
    path: '/projects/1',
    title: 'Uncategorized',
  },
  {
    id: 'trialDetails',
    needAuth: true,
    path: '/experiments/:experimentId/trials/:trialId/:tab?',
    title: 'Trial',
  },
  {
    id: 'trialDetails',
    needAuth: true,
    path: '/trials/:trialId/:tab?',
    title: 'Trial',
  },
  {
    id: 'experimentDetails',
    needAuth: true,
    path: '/experiments/:experimentId/:tab/:viz',
    title: 'Experiment',
  },
  {
    id: 'experimentDetails',
    needAuth: true,
    path: '/experiments/:experimentId/:tab',
    title: 'Experiment',
  },
  {
    id: 'experimentDetails',
    needAuth: true,
    path: '/experiments/:experimentId',
    title: 'Experiment',
  },
  {
    icon: 'logs',
    id: 'taskLogs',
    needAuth: true,
    path: '/:taskType/:taskId/logs',
    title: 'Task Logs',
  },
  {
    icon: 'tasks',
    id: 'taskList',
    needAuth: true,
    path: '/tasks',
    title: 'Tasks',
  },
  {
    id: 'modelVersionDetails',
    needAuth: true,
    path: '/models/:modelName/versions/:versionId',
    title: 'Version Details',
  },
  {
    id: 'modelDetails',
    needAuth: true,
    path: '/models/:modelName',
    title: 'Model Details',
  },
  {
    icon: 'model',
    id: 'models',
    needAuth: true,
    path: '/models',
    title: 'Model Registry',
  },
  {
    icon: 'cluster',
    id: 'cluster',
    needAuth: true,
    path: '/cluster/:tab?',
    title: 'Cluster',
  },
  {
    icon: 'cluster',
    id: 'clusters',
    needAuth: true,
    path: '/clusters/:tab?',
    title: 'Cluster',
  },
  {
    icon: 'logs',
    id: 'clusterLogs',
    needAuth: true,
    path: '/logs',
    title: 'Cluster Logs',
  },
  {
    id: 'wait',
    needAuth: true,
    path: '/wait/:taskType/:taskId',
  },
  {
    id: 'signIn',
    needAuth: false,
    path: '/login',
    title: 'Login',
  },
  {
    id: 'signOut',
    needAuth: false,
    path: '/logout',
    title: 'Logout',
  },
  {
    id: 'reload',
    needAuth: false,
    path: '/reload',
    title: 'Reload',
  },
  {
    id: 'jobs',
    needAuth: true,
    path: '/jobs',
    title: 'Job Queue',
  },
];

export default routes;
