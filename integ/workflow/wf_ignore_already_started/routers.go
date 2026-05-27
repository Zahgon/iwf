package wf_ignore_already_started

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

/**
 * This test workflow has 1 state, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- Waits on a timer
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType = "wf_ignore_already_started"
	State1       = "S1"
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

// fire after 10s

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
