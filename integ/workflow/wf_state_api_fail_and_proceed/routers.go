package wf_state_api_fail_and_proceed

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has one state, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- The state will fail and proceed to StateRecover which will gracefully complete workflow
 */
const (
	WorkflowType = "wf_state_api_fail_and_proceed"
	State1       = "S1"
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

// Bad Request response

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
