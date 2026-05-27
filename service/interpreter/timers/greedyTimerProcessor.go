package timers

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/interpreter/cont"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

type GreedyTimerProcessor struct {
	timerManager                    *timerScheduler
	stateExecutionCurrentTimerInfos map[string][]*service.TimerInfo
	staleSkipTimerSignals           []service.StaleSkipTimerSignal
	provider                        interfaces.WorkflowProvider
	logger                          interfaces.UnifiedLogger
}

func NewGreedyTimerProcessor(
	ctx interfaces.UnifiedContext,
	provider interfaces.WorkflowProvider,
	continueAsNewCounter *cont.ContinueAsNewCounter,
	staleSkipTimerSignals []service.StaleSkipTimerSignal,
) *GreedyTimerProcessor {
	_ = "STUB: not implemented"

	// start some single thread that manages pendingScheduling
	return nil
}

func (t *GreedyTimerProcessor) Dump() []service.StaleSkipTimerSignal {
	_ = "STUB: not implemented"
	return nil
}

func (t *GreedyTimerProcessor) GetTimerInfos() map[string][]*service.TimerInfo {
	_ = "STUB: not implemented"
	return nil
}

func (t *GreedyTimerProcessor) GetTimerStartedUnixTimestamps() []int64 {
	_ = "STUB: not implemented"
	return nil
}

// SkipTimer will attempt to skip a timer, return false if no valid timer found
func (t *GreedyTimerProcessor) SkipTimer(stateExeId, timerId string, timerIdx int) bool {
	_ = "STUB: not implemented"
	return false
}

// since we have checked it before sending signals, this should only happen in some vary rare cases for racing condition

func (t *GreedyTimerProcessor) RetryStaleSkipTimer() bool { _ = "STUB: not implemented"; return false }

// WaitForTimerFiredOrSkipped waits for timer completed(fired or skipped),
// return true when the timer is fired or skipped
// return false if the waitingCommands is canceled by cancelWaiting bool pointer(when the trigger type is completed, or continueAsNew)
func (t *GreedyTimerProcessor) WaitForTimerFiredOrSkipped(
	ctx interfaces.UnifiedContext, stateExeId string, timerIdx int, cancelWaiting *bool,
) service.InternalTimerStatus {
	_ = "STUB: not implemented"
	return *new(service.InternalTimerStatus)
}

// The waiting thread is later than the timer execState thread
// The execState thread got completed early and call RemovePendingTimersOfState to remove the timerInfos
// returning pending here

// This is trigger when one of the timers scheduled by the timerScheduler fires, scheduling a
//   new workflow task that evaluates the workflow's goroutines

// otherwise *cancelWaiting should return false to indicate that this timer isn't completed(fired or skipped)

// RemovePendingTimersOfState is for when a state is completed, remove all its pending pendingScheduling
func (t *GreedyTimerProcessor) RemovePendingTimersOfState(stateExeId string) {
	_ = "STUB: not implemented"
	return
}

func (t *GreedyTimerProcessor) AddTimers(
	stateExeId string, commands []iwfidl.TimerCommand, completedTimerCmds map[int]service.InternalTimerStatus,
) {
	_ = "STUB: not implemented"
	return
}
