package wf_force_fail

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has one state, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil method does nothing
 *      - Execute method will intentionally force-fail
 */
const (
	WorkflowType = "wf_force_fail"
	State1       = "S1"
)

type handler struct {
	invokeHistory sync.Map
}

var TestData = &iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("test-encoding"),
	Data:     iwfidl.PtrString("test-data"),
}

func NewHandler() common.WorkflowHandler {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandler)
}

func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Empty response

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Force fail

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
