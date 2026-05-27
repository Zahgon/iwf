package s3_init_data_attributes

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
)

/**
 * This test workflow has 2 states, testing S3 data attribute loading functionality.
 *
 * State1:
 *		- WaitUntil method loads and validates data attributes from S3
 *      - Execute method transitions to State2
 *
 * State2:
 *		- WaitUntil method does nothing
 *      - Execute method loads and validates data attributes from S3, then completes workflow
 */
const (
	WorkflowType      = "s3-init-data-attributes"
	State1            = "S1"
	State2            = "S2"
	TestDataAttrKey1  = "test-da-key1"
	TestDataAttrKey2  = "test-da-key2"
	TestDataAttrKey3  = "test-da-key3"
	LargeDataContent1 = "this_is_a_large_data_content_that_should_be_stored_in_s3_for_testing_purposes_with_more_than_10_characters"
	LargeDataContent2 = "another_large_data_content_for_second_attribute_that_exceeds_the_s3_threshold_for_external_storage_testing"
	SmallDataContent3 = "small"
)

var TestDataAttributeVal1 = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("\"" + LargeDataContent1 + "\""),
}

var TestDataAttributeVal2 = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("\"" + LargeDataContent2 + "\""),
}

var TestDataAttributeVal3 = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("\"" + SmallDataContent3 + "\""),
}

type handler struct {
	invokeHistory sync.Map
	invokeData    sync.Map
}

func NewHandler() *handler { _ = "STUB: not implemented"; return nil }

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Increment invoke count

// Store the state input for verification

// Validate that data attributes received match exactly the initial values provided at workflow start

// Increment invoke count

// State2 waitUntil doesn't need to check data attributes

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Increment invoke count

// Transition to State2

// Increment invoke count

// Validate that data attributes received match exactly the initial values provided at workflow start

// Complete workflow

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Merge both maps
