package persistence

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has three states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil method will update DA, SA, & SL
 * 		- Execute method will move to State2 with partially loaded data
 * State2:
 * 		- WaitUntil method will store attribute data
 * 		- Execute method will move to State3 with partially loaded data
 * State3:
 * 		- WaitUntil method performs some attribute checks
 * 		- Execute method performs checks on the attribute data and then gracefully completes the workflow
 */
const (
	WorkflowType          = "persistence"
	State1                = "S1"
	State2                = "S2"
	State3                = "S3"
	TestDataAttributeKey  = "test-data-attribute"
	TestDataAttributeKey2 = "test-data-attribute-2"
	TestStateLocalKey     = "test-state-local"

	TestSearchAttributeKeywordKey    = "CustomKeywordField"
	TestSearchAttributeKeywordValue1 = "keyword-value1"
	TestSearchAttributeKeywordValue2 = "keyword-value2"

	TestSearchAttributeKeywordArrayKey = "CustomKeywordArrayField"
	TestSearchAttributeIntKey          = "CustomIntField"
	TestSearchAttributeBoolKey         = "CustomBoolField"
	TestSearchAttributeDoubleKey       = "CustomDoubleField"
	TestSearchAttributeDatetimeKey     = "CustomDatetimeField"
	TestSearchAttributeTextKey         = "CustomStringField"
	TestSearchAttributeIntValue1       = 1
	TestSearchAttributeIntValue2       = 2
)

var TestDataAttributeVal1 = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-data-attribute-value1"),
}

var TestDataAttributeVal2 = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-data-attribute-value2"),
}

var testStateLocalVal = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-state-local-value"),
}

type handler struct {
	invokeHistory sync.Map
	invokeData    sync.Map
}

func NewHandler() common.WorkflowHandler {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandler)
}

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Go to the decide methods after updating DA, SA, & SL

// Determine how many keywords and ints are found in the search attributes

// Determine if the attribute is found in the request

// Go straight to the decide methods without any commands

// Determine if the INT attribute is found in the request

// Go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Determine how many keywords and ints are found in the search attributes

// Determine how many query attributes are found

// Determine if local attribute is found

// Move to state 2 with set options after updating values

// Determine how many keywords and ints are found in the search attributes

// Determine how many query attributes are found

// Move to state 3 after with set options

// Determine if the INT attribute is found in the request

// Determine how many query attributes are found

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
