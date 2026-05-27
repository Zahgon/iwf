package wait_until_search_attributes

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
 *		- Waits on nothing. Will execute momentarily
 *      - Execute method will go to State2
 * State2:
 *		- Waits on nothing. Will execute momentarily
 *		- Skips wait until
 *      - 10-second delay is added before executing state
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType = "wait_until_search_attributes"
	State1       = "S1"
	State2       = "S2"

	TestSearchAttributeExecutingStateIdsKey = "IwfExecutingStateIds"
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

// Go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to State 2, skipping wait until

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
