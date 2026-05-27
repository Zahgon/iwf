package anycommandcombination

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
 *		- WaitUntil will fail its first attempt and then retry which will proceed when a combination is completed
 *      - Execute method will invoke the combination and move the State2
 * State2:
 *		- WaitUntil will fail its first attempt and then retry which will proceed when a combination is completed
 *      - Execute method will invoke the combination and gracefully complete workflow
 */
const (
	WorkflowType     = "any_command_combination"
	State1           = "S1"
	State2           = "S2"
	TimerId1         = "test-timer-1"
	SignalNameAndId1 = "test-signal-name1"
	SignalNameAndId2 = "test-signal-name2"
	SignalNameAndId3 = "test-signal-name3"
)

type handler struct {
	invokeHistory sync.Map
	invokeData    sync.Map
	//we want to confirm that the interpreter workflow activity will fail when the commandId is empty with ANY_COMMAND_COMBINATION_COMPLETED
	hasS1RetriedForInvalidCommandId bool
	hasS2RetriedForInvalidCommandId bool
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

// one year later

// one year later

// If the state has already retried an invalid command, proceed on combination completed

// wait for two SignalNameAndId1

// If the state has not already retried an invalid command, return invalid trigger signals, which will fail
// and cause a retry

// If the state has already retried an invalid command, return signals and completion metrics

// If the state has not already retried an invalid command, return invalid trigger signals, which will fail
// and cause a retry

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Trigger signals and move to State 2

// Trigger data and move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
