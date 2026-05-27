package deadend

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has 3 states, using REST controller to implement the workflow directly.
 *
 * RPCWriteData:
 *		- WaitUntil will upsert data attributes
 * RPCTriggerState:
 *		- WaitUntil will move to State1
 * State1:
 *		- WaitUntil is skipped
 *      - Execute method will put the state into a dead-end.
 */
const (
	WorkflowType    = "deadend"
	RPCTriggerState = "test-RPCTriggerState"
	RPCWriteData    = "RPCWriteData"

	State1 = "S1"
)

type handler struct {
	invokeHistory sync.Map
}

func NewHandler() common.WorkflowHandlerWithRpc {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandlerWithRpc)
}

func (h *handler) ApiV1WorkflowWorkerRpc(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to State 1

// Upsert data attributes

// ApiV1WorkflowStateStart - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to the dead-end state

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
