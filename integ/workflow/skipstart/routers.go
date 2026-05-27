package skipstart

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

/**
 * This test workflow has 2 states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- Wait until is skipped.
 *      - Execute method will go to State2
 * State2:
 *		- Wait until is skipped.
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType = "skipstart"
	State1       = "S1"
	State2       = "S2"
)

type handler struct {
	invokeHistory sync.Map
}

func NewHandler() *handler { _ = "STUB: not implemented"; return nil }

// ApiV1WorkflowStateStart - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to State 2 with the provided input & options

// Move to completion with the provided input

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
