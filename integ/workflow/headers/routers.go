package headers

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

/**
 * This test workflow has 1 state, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil method does nothing
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType = "headers"
	State1       = "S1"

	TestHeaderKey   = "integration-test-header"
	TestHeaderValue = "integration-test-value"
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

// Basic workflow to go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
