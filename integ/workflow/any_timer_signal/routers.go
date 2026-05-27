package anytimersignal

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has 2 states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil will wait for the signals to trigger.
 *      - Execute method will trigger signals and retry State1 once, then trigger signals and move the State2
 * State2:
 *		- Waits on nothing. Will execute momentarily
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType = "any_timer_signal"
	State1       = "S1"
	State2       = "S2"
	SignalName   = "test-signal-name"
)

type handler struct {
	invokeHistory sync.Map
	invokeData    sync.Map
}

func NewHandler() common.WorkflowHandler {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandler)
}

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Fire timer after 1s on first start attempt

// fire after 1s

// Go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// On first State 1 attempt, trigger signals and stay on the first state

// After the first State 1 attempt, trigger signals and move to next state

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
