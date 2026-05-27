package timer

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
 *		- Has 3 timers (10s, 1d, 1y) before executing state
 *      - Execute method will go to State2
 * State2:
 *		- Waits on nothing. Will execute momentarily
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType = "timer"
	State1       = "S1"
	State2       = "S2"
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

// Proceed after 3 timers complete

// fire after 10s

// fire after one day

// fire after one year

// Go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to State 2

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
