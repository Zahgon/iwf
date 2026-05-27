package wf_state_options_search_attributes_loading

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has five states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil method does nothing
 * 		- Execute method sets values for all Search Attributes used in this test
 * 			- CustomKeywordField
 * 			- CustomStringField
 *			- CustomBoolField
 * State2:
 * 		- State Options contains WaitUntilApiSearchAttributesLoadingPolicy
 * 		- WaitUntil method asserts that expected SearchAttributes are loaded
 * 		- Execute method asserts that no SearchAttributes are loaded
 * State3:
 * 		- State Options contains WaitUntilApiSearchAttributesLoadingPolicy
 * 		- WaitUntil method asserts that no SearchAttributes are loaded
 * 		- Execute method asserts that expected SearchAttributes are loaded
 * State4:
 * 		- State Options contains SearchAttributesLoadingPolicy
 * 		- WaitUntil method asserts that expected SearchAttributes are loaded
 * 		- Execute method asserts that expected SearchAttributes are loaded
 * State5:
 * 		- State Options contains AdSearchAttributesLoadingPolicy and WaitUntilApiSearchAttributesLoadingPolicy
 * 		- WaitUntil method asserts that WaitUntilApiSearchAttributesLoadingPolicy are loaded
 * 		- Execute method asserts that SearchAttributesLoadingPolicy are loaded
 */
const (
	WorkflowType = "state_options_search_attributes_loading"
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

// Move to State 2 with the provided options & input after updating data attributes

func getState2DecideResponse(req iwfidl.WorkflowStateDecideRequest) iwfidl.WorkflowStateDecideResponse {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowStateDecideResponse)
}

// Move to State 3 with the provided options & input

func getState3DecideResponse(req iwfidl.WorkflowStateDecideRequest) iwfidl.WorkflowStateDecideResponse {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowStateDecideResponse)
}

// Move to State 4 with the provided options & input

func getState4DecideResponse(req iwfidl.WorkflowStateDecideRequest) iwfidl.WorkflowStateDecideResponse {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowStateDecideResponse)
}

// Move to State 5 with the provided options & input

func getState5DecideResponse() iwfidl.WorkflowStateDecideResponse {
	_ = "STUB: not implemented"
	// Move to completion
	return *new(iwfidl.WorkflowStateDecideResponse)
}

func verifyEmptySearchAttributes(t *testing.T, searchAttributes []iwfidl.SearchAttribute) {
	_ = "STUB: not implemented"
	return
}

func verifyLoadedSearchAttributes(t *testing.T, stateId string, method string, searchAttributes []iwfidl.SearchAttribute, loadingType iwfidl.PersistenceLoadingType) {
	_ = "STUB: not implemented"
	return
}

func getUpsertSearchAttributes() []iwfidl.SearchAttribute { _ = "STUB: not implemented"; return nil }

func getExpectedSearchAttributes(stateId string, method string, loadingType iwfidl.PersistenceLoadingType) []iwfidl.SearchAttribute {
	_ = "STUB: not implemented"
	return nil
}

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
