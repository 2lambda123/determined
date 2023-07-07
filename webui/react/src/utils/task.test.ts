import { CommandState, CommandTask, CommandType, ExperimentTask, id, RunState, Task } from 'types';

import { canBeOpened, isExperimentTask } from './task';

const SampleTask: Task = { id: '', name: '', resourcePool: '', startTime: '' };
const SampleExperimentTask: ExperimentTask = {
  ...SampleTask,
  archived: false,
  parentArchived: false,
  projectId: id(0),
  resourcePool: '',
  state: 'ACTIVE' as RunState,
  userId: id(345),
  username: '',
  workspaceId: id(0),
};
const SampleCommandTask: CommandTask = {
  ...SampleTask,
  resourcePool: '',
  state: 'PENDING' as CommandState,
  type: 'COMMAND' as CommandType,
  userId: id(345),
  workspaceId: id(0),
};

describe('isExperimentTask', () => {
  it('Experiment Task', () => {
    expect(isExperimentTask(SampleExperimentTask)).toStrictEqual(true);
  });
  it('Command Task', () => {
    expect(isExperimentTask(SampleCommandTask)).toStrictEqual(false);
  });
});

describe('canBeOpened', () => {
  it('Experiment Task', () => {
    expect(canBeOpened(SampleExperimentTask)).toStrictEqual(true);
  });
  it('Terminated Command Task', () => {
    expect(
      canBeOpened({ ...SampleCommandTask, state: 'TERMINATED' as CommandState }),
    ).toStrictEqual(false);
  });
  it('Command Task without service address', () => {
    expect(canBeOpened(SampleCommandTask)).toStrictEqual(false);
  });
  it('Command Task with service address', () => {
    expect(canBeOpened({ ...SampleCommandTask, serviceAddress: 'test' })).toStrictEqual(true);
  });
});
