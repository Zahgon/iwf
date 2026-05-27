package wf_state_options_data_attributes_loading

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has four states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil method does nothing
 * 		- Execute method creates all Data Attributes keys that will be used in this test
 * 			- da_wait_until1
 * 			- da_execute1
 *			- da_other_key
 * State2:
 * 		- State Options contains WaitUntilApiDataAttributesLoadingPolicy
 * 		- WaitUntil method asserts that expected DataAttributes are loaded
 * 		- Execute method asserts that no DataAttributes are loaded
 * State3:
 * 		- State Options contains ExecuteApiDataAttributesLoadingPolicy
 * 		- WaitUntil method asserts that no DataAttributes are loaded
 * 		- Execute method asserts that expected DataAttributes are loaded
 * State4:
 * 		- State Options contains DataAttributesLoadingPolicy
 * 		- WaitUntil method asserts that expected DataAttributes are loaded
 * 		- Execute method asserts that expected DataAttributes are loaded
 * State5:
 * 		- State Options contains DataAttributesLoadingPolicy and WaitUntilApiDataAttributesLoadingPolicy
 * 		- WaitUntil method asserts that WaitUntilApiDataAttributesLoadingPolicy are loaded
 * 		- Execute method asserts that DataAttributesLoadingPolicy are loaded
 */
const (
	WorkflowType = "state_options_data_attributes_loading"
	State1       = "S1"
	State2       = "S2"
	State3       = "S3"
	State4       = "S4"
	State5       = "S5"
)

type handler struct {
	invokeHistory sync.Map
}

func NewHandler() common.WorkflowHandlerWithRpc {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandlerWithRpc)
}

func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) ApiV1WorkflowWorkerRpc(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func getState1DecideResponse(req iwfidl.WorkflowStateDecideRequest) iwfidl.WorkflowStateDecideResponse {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowStateDecideResponse)
}

// Move to State 2 with provided options & input after updating data attributes

func getState2DecideResponse(req iwfidl.WorkflowStateDecideRequest) iwfidl.WorkflowStateDecideResponse {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowStateDecideResponse)
}

// Move to State 3 with provided options & input

func getState3DecideResponse(req iwfidl.WorkflowStateDecideRequest) iwfidl.WorkflowStateDecideResponse {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowStateDecideResponse)
}

// Move to State 4 with provided options & input

func getState4DecideResponse(req iwfidl.WorkflowStateDecideRequest) iwfidl.WorkflowStateDecideResponse {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowStateDecideResponse)
}

// Move to State 5 with provided options & input

func getState5DecideResponse() iwfidl.WorkflowStateDecideResponse {
	_ = "STUB: not implemented"
	// Move to completion
	return *new(iwfidl.WorkflowStateDecideResponse)
}

func verifyEmptyDataAttributes(t *testing.T, dataAttributes []iwfidl.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func verifyLoadedDataAttributes(t *testing.T, stateId string, method string, dataAttributes []iwfidl.KeyValue, loadingType iwfidl.PersistenceLoadingType) {
	_ = "STUB: not implemented"
	return
}

func getUpsertDataAttributes() []iwfidl.KeyValue { _ = "STUB: not implemented"; return nil }

func getExpectedDataAttributes(stateId string, method string, loadingType iwfidl.PersistenceLoadingType) []iwfidl.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
