package reset

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

/**
* This test workflow has 2 states, using REST controller to implement the workflow directly.
* State1:
*       - No WaitUntil
*       - Execute moves to State2
* State2:
* 		- No WaitUntil
*       - Execute loops through state2 5 times, then gracefully completes the workflow.
* This test is used for testing reset by state id and state execution id without WaitUntil
 */
const (
	WorkflowType = "reset"
	State1       = "S1"
	State2       = "S2"
)

type handler struct {
	invokeHistory sync.Map
}

func NewHandler() *handler { _ = "STUB: not implemented"; return nil }

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// go to S2

//Skipping wait until for 1st execution of state2

//Skipping wait until for all executions of state2 after the 1st execution.

// go to complete

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
