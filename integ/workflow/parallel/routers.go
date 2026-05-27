package parallel

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has eight states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil method does nothing
 * 		- Execute method delays 1s then moves to State11, State12, & State13
 * State11:
 *		- WaitUntil method does nothing
 * 		- Execute method delays 2s then moves to State111 & State112
 * State12:
 *		- WaitUntil method does nothing
 * 		- Execute method delays 2s then moves to State121 & State122
 * State13:
 *		- WaitUntil method does nothing
 *      - Execute method will delay 1s then gracefully complete workflow
 * State111:
 *		- WaitUntil method does nothing
 *      - Execute method will gracefully complete workflow
 * State112:
 *		- WaitUntil method does nothing
 *      - Execute method will gracefully complete workflow
 * State121:
 *		- WaitUntil method does nothing
 *      - Execute method will gracefully complete workflow
 * State122:
 *		- WaitUntil method does nothing
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType = "parallel"
	State1       = "S1"
	State11      = "S11"
	State12      = "S12"
	State13      = "S13"
	State111     = "S111"
	State112     = "S112"
	State121     = "S121"
	State122     = "S122"
)

type handler struct {
	invokeHistory sync.Map
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

// Go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Cause graceful complete to wait

// Move to 3 states (which will all move to this decide method without commands)

// Cause graceful complete to wait

// Move to 2 states (which will all move to this decide method without commands)

// Cause graceful complete to wait

// Move to 2 states (which will all move to this decide method without commands)

// Cause graceful complete to wait

// Move to completion after updating the state input

// Move to completion after updating the state input

// Fail workflow due to unknown or unexpected state

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
