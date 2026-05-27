package wait_until_search_attributes_optimization

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has 7 states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- Waits one second before executing
 *      - Execute method will loop back to State1 five times; then execute method will go to State2
 * State2:
 *		- First execution: loops back to State2 + goes to State3
 *      - Second execution (after 1 second): goes to State3 and State4
 * State3:
 *		- Waits 8 seconds
 *      - Execute method will gracefully complete workflow
 * State4:
 *		- Waits on command trigger (signal)
 *      - Execute method will go to State5
 * State5:
 *		- Skips waitUntil and executes momentarily
 *      - Execute method will go to State6 and State7
 * State6:
 *		- Waits 4 seconds
 *      - Execute method will gracefully complete workflow
 * State7:
 *		- Skips waitUntil and executes momentarily
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType = "wait_until_search_optimization"
	State1       = "S1"
	State2       = "S2"
	State3       = "S3"
	State4       = "S4"
	State5       = "S5"
	State6       = "S6"
	State7       = "S7"

	SignalName = "test-signal"
)

type handler struct {
	invokeHistory sync.Map
}

func NewHandler() common.WorkflowHandler {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandler)
}

func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Go straight to the decide methods without any commands

// Proceed after signal is received

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to State 2

// Repeat State 1 (5 times)

// Move to State 3 & 4

// Repeat State 2 and Move to State 3

// Move to Completion

// Move to State 5, skipping wait until

// Move to State 6 and State 7 skipping wait until for 7

// Move to completion

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
