package timers

import (
	"github.com/indeedeng/iwf/service/interpreter/interfaces"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
)

type SimpleTimerProcessor struct {
	stateExecutionCurrentTimerInfos map[string][]*service.TimerInfo
	awaitingTimers                  []int64
	staleSkipTimerSignals           []service.StaleSkipTimerSignal
	provider                        interfaces.WorkflowProvider
	logger                          interfaces.UnifiedLogger
}

func NewSimpleTimerProcessor(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, staleSkipTimerSignals []service.StaleSkipTimerSignal,
) *SimpleTimerProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (t *SimpleTimerProcessor) Dump() []service.StaleSkipTimerSignal {
	_ = "STUB: not implemented"
	return nil
}

func (t *SimpleTimerProcessor) GetTimerInfos() map[string][]*service.TimerInfo {
	_ = "STUB: not implemented"
	return nil
}

func (t *SimpleTimerProcessor) GetTimerStartedUnixTimestamps() []int64 {
	_ = "STUB: not implemented"
	return nil

	// SkipTimer will attempt to skip a timer, return false if no valid timer found
}

func (t *SimpleTimerProcessor) SkipTimer(stateExeId, timerId string, timerIdx int) bool {
	_ = "STUB: not implemented"
	return false
}

// since we have checked it before sending signals, this should only happen in some vary rare cases for racing condition

func (t *SimpleTimerProcessor) RetryStaleSkipTimer() bool { _ = "STUB: not implemented"; return false }

// WaitForTimerFiredOrSkipped waits for timer completed(fired or skipped),
// return true when the timer is fired or skipped
// return false if the waitingCommands is canceled by cancelWaiting bool pointer(when the trigger type is completed, or continueAsNew)
func (t *SimpleTimerProcessor) WaitForTimerFiredOrSkipped(
	ctx interfaces.UnifiedContext, stateExeId string, timerIdx int, cancelWaiting *bool,
) service.InternalTimerStatus {
	_ = "STUB: not implemented"
	return *new(service.InternalTimerStatus)
}

// The waiting thread is later than the timer execState thread
// The execState thread got completed early and call RemovePendingTimersOfState to remove the timerInfos
// returning pending here

// otherwise *cancelWaiting should return false to indicate that this timer isn't completed(fired or skipped)

func removeSingleTime(timers []int64, at int64) []int64 { _ = "STUB: not implemented"; return nil }

// RemovePendingTimersOfState is for when a state is completed, remove all its pending timers
func (t *SimpleTimerProcessor) RemovePendingTimersOfState(stateExeId string) {
	_ = "STUB: not implemented"
	return
}

func (t *SimpleTimerProcessor) AddTimers(
	stateExeId string, commands []iwfidl.TimerCommand, completedTimerCmds map[int]service.InternalTimerStatus,
) {
	_ = "STUB: not implemented"
	return
}
