package interpreter

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/interpreter/config"
	"github.com/indeedeng/iwf/service/interpreter/cont"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

type StateExecutionCounter struct {
	ctx                  interfaces.UnifiedContext
	provider             interfaces.WorkflowProvider
	configer             *config.WorkflowConfiger
	globalVersioner      *GlobalVersioner
	continueAsNewCounter *cont.ContinueAsNewCounter

	stateIdCompletedCounts          map[string]int
	stateIdStartedCounts            map[string]int // For creating stateExecutionId: count the stateId for how many times that have been executed
	stateIdCurrentlyExecutingCounts map[string]int // For system search attribute IwfExecutingStateId: keep counting the stateIds that are executing based on the ExecutingStateIdMode
	totalCurrentlyExecutingCount    int            // For "dead ends": count the total pending states
}

func NewStateExecutionCounter(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, globalVersioner *GlobalVersioner,
	configer *config.WorkflowConfiger, continueAsNewCounter *cont.ContinueAsNewCounter,
) *StateExecutionCounter {
	_ = "STUB: not implemented"
	return nil
}

func RebuildStateExecutionCounter(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, globalVersioner *GlobalVersioner,
	stateIdStartedCounts map[string]int, stateIdCurrentlyExecutingCounts map[string]int,
	totalCurrentlyExecutingCount int,
	configer *config.WorkflowConfiger, continueAsNewCounter *cont.ContinueAsNewCounter,
) *StateExecutionCounter {
	_ = "STUB: not implemented"
	return nil
}

func (e *StateExecutionCounter) Dump() service.StateExecutionCounterInfo {
	_ = "STUB: not implemented"
	return *new(service.StateExecutionCounterInfo)
}

func (e *StateExecutionCounter) CreateNextExecutionId(stateId string) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *StateExecutionCounter) MarkStateIdExecutingIfNotYet(stateReqs []StateRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// do nothing

func (e *StateExecutionCounter) increaseStateIdCurrentlyExecutingCounts(s iwfidl.StateMovement) bool {
	_ = "STUB: not implemented"
	return false
}

// first time the stateId show up

func (e *StateExecutionCounter) MarkStateExecutionCompleted(currentState iwfidl.StateMovement, nextStates []iwfidl.StateMovement) error {
	_ = "STUB: not implemented"
	return nil
}

func determineIfShouldSkipRefreshOnCompleted(nextStates []iwfidl.StateMovement, enabledForAll bool) bool {
	_ = "STUB: not implemented"
	return false
}

// s is not a ValidClosingWorkflowStateId

func (e *StateExecutionCounter) decreaseStateIdCurrentlyExecutingCounts(state iwfidl.StateMovement) {
	_ = "STUB: not implemented"
	return
}

func (e *StateExecutionCounter) GetTotalCurrentlyExecutingCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (e *StateExecutionCounter) refreshIwfExecutingStateIdSearchAttribute() error {
	_ = "STUB: not implemented"
	// Optimization: don't upsert SAs if currentSAsValues == stateIdCurrentlyExecutingCounts keys
	return nil
}

// we don't clear search attributes because there are only two possible cases:
// 1. there will be another stateId being upsert right after this. So this will avoid calling the upsertSA twice
// 2. there will not be another stateId being upsert. Then this will be cleared before the workflow is closed.
// see workflowImpl.go to call ClearExecutingStateIdsSearchAttributeFinally at the end

// ClearExecutingStateIdsSearchAttributeFinally should only be called at the end of workflow
func (e *StateExecutionCounter) ClearExecutingStateIdsSearchAttributeFinally() {
	_ = "STUB: not implemented"
	return
}
