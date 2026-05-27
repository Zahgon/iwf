package interpreter

import (
	"time"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

type ContinueAsNewer struct {
	provider interfaces.WorkflowProvider

	StateExecutionToResumeMap map[string]service.StateExecutionResumeInfo // stateExeId to StateExecutionResumeInfo
	inflightUpdateOperations  int

	stateRequestQueue     *StateRequestQueue
	interStateChannel     *InternalChannel
	stateExecutionCounter *StateExecutionCounter
	persistenceManager    *PersistenceManager
	signalReceiver        *SignalReceiver
	outputCollector       *OutputCollector
	timerProcessor        interfaces.TimerProcessor
}

func NewContinueAsNewer(
	provider interfaces.WorkflowProvider,
	interStateChannel *InternalChannel, signalReceiver *SignalReceiver, stateExecutionCounter *StateExecutionCounter,
	persistenceManager *PersistenceManager, stateRequestQueue *StateRequestQueue, collector *OutputCollector,
	timerProcessor interfaces.TimerProcessor,
) *ContinueAsNewer {
	_ = "STUB: not implemented"
	return nil
}

func LoadInternalsFromPreviousRun(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, previousRunId string, continueAsNewPageSizeInBytes int32,
) (*service.ContinueAsNewDumpResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reset to start from beginning

func (c *ContinueAsNewer) GetSnapshot() service.ContinueAsNewDumpResponse {
	_ = "STUB: not implemented"
	return *new(service.ContinueAsNewDumpResponse)
}

func (c *ContinueAsNewer) SetQueryHandlersForContinueAsNew(ctx interfaces.UnifiedContext) error {
	_ = "STUB: not implemented"
	return nil
}

// return the current page of the whole snapshot

func (c *ContinueAsNewer) AddPotentialStateExecutionToResume(
	stateExecutionId string, state iwfidl.StateMovement, stateExecLocals []iwfidl.KeyValue,
	commandRequest iwfidl.CommandRequest,
	completedTimerCommands map[int]service.InternalTimerStatus,
	completedSignalCommands, completedInterStateChannelCommands map[int]*iwfidl.EncodedObject,
) {
	_ = "STUB: not implemented"
	return
}

func (c *ContinueAsNewer) HasAnyStateExecutionToResume() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ContinueAsNewer) RemoveStateExecutionToResume(stateExecutionId string) {
	_ = "STUB: not implemented"
	return
}

func (c *ContinueAsNewer) DrainThreads(ctx interfaces.UnifiedContext) error {
	_ = "STUB: not implemented"
	// TODO: add metric for before and after Await to monitor stuck
	// NOTE: consider using AwaitWithTimeout to get an alert when workflow stuck due to a bug in the draining logic for continueAsNew
	return nil
}

func (c *ContinueAsNewer) IncreaseInflightOperation() { _ = "STUB: not implemented"; return }

func (c *ContinueAsNewer) DecreaseInflightOperation() { _ = "STUB: not implemented"; return }

// if the DrainAllSignalsAndThreads await is being called more than a few times and cannot get through,
// there is likely something wrong in the continueAsNew logic (unless state API is stuck)
// the key is runId, the value is how many times it has been called in this worker
// Using this in memory counter sot hat we don't have to use AwaitWithTimeout which will consume a timer
// TODO add TTL support because we don't have to keep the value in memory forever(likely a few hours or a day is enough)
var inMemoryContinueAsNewMonitor = make(map[string]time.Time)

const warnThreshold = time.Second * 5
const errThreshold = time.Second * 15

func (c *ContinueAsNewer) allThreadsDrained(ctx interfaces.UnifiedContext) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO using a flag to control this debugging info
