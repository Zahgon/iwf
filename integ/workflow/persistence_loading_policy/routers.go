package persistence_loading_policy

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has two states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil skipped
 * 		- Execute method verifies the loaded attributes then moves to a dead-end.
 * State2:
 * 		- WaitUntil method verifies the loaded attributes
 * 		- Execute method verifies the loaded attributes then gracefully completes the workflow
 */
const (
	WorkflowType = "persistence_loading_policy"
	State1       = "S1"
	State2       = "S2"
)

type handler struct {
	invokeHistory sync.Map
}

func NewHandler() common.WorkflowHandlerWithRpc {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandlerWithRpc)
}

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Dynamically get the loadingType from input

// Go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Dynamically get the loadingType from input

// Set search attributes and data attributes in State1

// Move to dead-end (state 1) or completion (state 2)

func (h *handler) ApiV1WorkflowWorkerRpc(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// dynamically get the loadingType from input

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

func verifyLoadedAttributes(
	t *testing.T,
	searchAttributes []iwfidl.SearchAttribute,
	dataAttributes []iwfidl.KeyValue,
	loadingType iwfidl.PersistenceLoadingType) {
	_ = "STUB: not implemented"
	return
}

// use ElementsMatch so that the order won't be a problem.
// Internally the SAs are stored as a map and as a result, Golang return it without ordering guarantee

func getStateDecision(nextStateId string, loadingTypeFromInput iwfidl.EncodedObject, loadingType iwfidl.PersistenceLoadingType) *iwfidl.StateDecision {
	_ = "STUB: not implemented"
	return nil
}
