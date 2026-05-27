package timers

import (
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/interpreter/cont"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

type timerScheduler struct {
	// Timers requested by the workflow in desc order
	pendingScheduling []*service.TimerInfo
	// timers created through the workflow provider that are going to fire in desc order
	providerScheduledTimerUnixTs []int64
}

func (t *timerScheduler) addTimer(toAdd *service.TimerInfo) { _ = "STUB: not implemented"; return }

// don't want dupes. Makes remove simpler

func (t *timerScheduler) removeTimer(toRemove *service.TimerInfo) {
	_ = "STUB: not implemented"
	return
}

func (t *timerScheduler) pruneToNextTimer(pruneTo int64) *service.TimerInfo {
	_ = "STUB: not implemented"
	return nil
}

// If index is 0, it means all times are in the past

// If index is 0, it means all timers are pruned

func startGreedyTimerScheduler(
	ctx interfaces.UnifiedContext,
	provider interfaces.WorkflowProvider,
	continueAsNewCounter *cont.ContinueAsNewCounter) *timerScheduler {
	_ = "STUB: not implemented"
	return nil
}

// Expecting CanceledError if the ctx is canceled

// This will create a new timer but not yield the goroutines awaiting the timer firing.
// This works since when a timer fires, a new workflow task is created with the expectation that
//   there is a goroutines awaiting some condition(some time has past) to continue,
//   see WaitForTimerFiredOrSkipped.
