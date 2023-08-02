//nolint:exhaustivestruct
package internal

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/determined-ai/determined/master/internal/rm/actorrm"
	"github.com/determined-ai/determined/master/pkg/actor/actors"
	"github.com/determined-ai/determined/master/pkg/ssh"

	"github.com/pkg/errors"

	"github.com/determined-ai/determined/master/internal/task"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/internal/mocks"
	"github.com/determined-ai/determined/master/internal/mocks/allocationmocks"
	"github.com/determined-ai/determined/master/pkg/actor"
	"github.com/determined-ai/determined/master/pkg/etc"
	detLogger "github.com/determined-ai/determined/master/pkg/logger"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/master/pkg/ptrs"
	"github.com/determined-ai/determined/master/pkg/schemas"
	"github.com/determined-ai/determined/master/pkg/schemas/expconf"
	"github.com/determined-ai/determined/master/pkg/searcher"
	"github.com/determined-ai/determined/master/pkg/tasks"
)

func TestTrial(t *testing.T) {
	system, db, rID, tr, self, alloc := setup(t)

	// Pre-scheduled stage.
	db.On("UpdateTrialRunID", 0, 1).Return(nil)
	db.On("LatestCheckpointForTrial", 0).Return(&model.Checkpoint{}, nil)
	require.NoError(t, system.Ask(self,
		model.StateWithReason{State: model.ActiveState}).Error())
	require.NoError(t, system.Ask(self, trialSearcherState{
		Create: searcher.Create{RequestID: rID},
		Op: searcher.ValidateAfter{
			RequestID: rID,
			Length:    10,
		},
		Complete: false,
		Closed:   true,
	}).Error())
	require.True(t, alloc.AssertExpectations(t))
	require.NotNil(t, tr.allocationID)
	require.True(t, db.AssertExpectations(t))

	// Running stage.
	db.On("UpdateTrial", 0, model.StoppingCompletedState).Return(nil)
	require.NoError(t, system.Ask(self, trialSearcherState{
		Create: searcher.Create{RequestID: rID},
		Op: searcher.ValidateAfter{
			RequestID: rID,
			Length:    10,
		},
		Complete: true,
		Closed:   true,
	}).Error())
	require.True(t, db.AssertExpectations(t))

	// Terminating stage.
	db.On("UpdateTrial", 0, model.CompletedState).Return(nil)
	system.Tell(self, &task.AllocationExited{})
	require.NoError(t, self.AwaitTermination())
	require.True(t, model.TerminalStates[tr.state])
	require.True(t, db.AssertExpectations(t))
}

func TestTrialRestarts(t *testing.T) {
	system, db, rID, tr, self, _ := setup(t)

	// Pre-scheduled stage.
	db.On("UpdateTrialRunID", 0, 1).Return(nil)
	db.On("LatestCheckpointForTrial", 0).Return(&model.Checkpoint{}, nil)
	require.NoError(t, system.Ask(self,
		model.StateWithReason{State: model.ActiveState}).Error())
	require.NoError(t, system.Ask(self, trialSearcherState{
		Create: searcher.Create{RequestID: rID},
		Op: searcher.ValidateAfter{
			RequestID: rID,
			Length:    10,
		},
		Complete: false,
		Closed:   true,
	}).Error())
	require.True(t, db.AssertExpectations(t))

	for i := 0; i <= tr.config.MaxRestarts(); i++ {
		require.NotNil(t, tr.allocationID)
		require.Equal(t, i, tr.restarts)

		db.On("UpdateTrialRestarts", 0, i+1).Return(nil)
		if i == tr.config.MaxRestarts() {
			db.On("UpdateTrial", 0, model.ErrorState).Return(nil)
		} else {
			// For the next go-around, when we update trial run ID.
			db.On("UpdateTrialRunID", 0, i+2).Return(nil)
		}

		system.Tell(self, &task.AllocationExited{Err: errors.New("bad stuff went down")})
		system.Ask(self, actor.Ping{}).Get() // sync

		require.True(t, db.AssertExpectations(t))
	}
	require.NoError(t, self.AwaitTermination())
	require.True(t, model.TerminalStates[tr.state])
}

func setup(t *testing.T) (
	*actor.System,
	*mocks.DB,
	model.RequestID,
	*trial,
	*actor.Ref,
	*allocationmocks.AllocationService,
) {
	require.NoError(t, etc.SetRootPath("../static/srv"))
	system := actor.NewSystem("system")

	// mock resource manager.
	rmActor := actors.MockActor{Responses: map[string]*actors.MockResponse{}}
	rmImpl := actorrm.Wrap(system.MustActorOf(actor.Addr("rm"), &rmActor))

	// mock allocation service
	var as allocationmocks.AllocationService
	task.DefaultService = &as
	as.On(
		"StartAllocation", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything, mock.Anything,
	).Return()

	// mock db.
	db := &mocks.DB{}
	db.On("AddTrial", mock.Anything).Return(nil)
	db.On("AddTask", mock.Anything).Return(nil)
	db.On("UpdateTrial", mock.Anything, model.ActiveState).Return(nil)

	// instantiate the trial
	rID := model.NewRequestID(rand.Reader)
	taskID := model.TaskID(fmt.Sprintf("%s-%s", model.TaskTypeTrial, rID))
	tr, _ := newTrial(
		detLogger.Context{},
		taskID,
		model.JobID("1"),
		time.Now(),
		1,
		model.PausedState,
		trialSearcherState{Create: searcher.Create{RequestID: rID}, Complete: true},
		rmImpl,
		db,
		schemas.WithDefaults(expconf.ExperimentConfig{
			RawCheckpointStorage: &expconf.CheckpointStorageConfigV0{
				RawSharedFSConfig: &expconf.SharedFSConfig{
					RawHostPath:      ptrs.Ptr("/tmp"),
					RawContainerPath: ptrs.Ptr("determined-sharedfs"),
				},
			},
		}),
		&model.Checkpoint{},
		&tasks.TaskSpec{
			AgentUserGroup: &model.AgentUserGroup{},
			SSHRsaSize:     1024,
		},
		ssh.PrivateAndPublicKeys{},
		false,
		nil, false, nil, nil,
	)
	// self := system.MustActorOf(actor.Addr("trial"), tr)
	return system, db, rID, tr, nil, &as
}
