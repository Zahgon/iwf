package s3GetSetDataAttributes

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
)

/**
 * Test workflow for S3 external storage with get/set data attributes APIs.
 * Tests both small data (stays in Temporal) and large data (goes to S3).
 *
 * State1:
 *   - Simple workflow that waits and completes
 *
 * The main testing is done via direct API calls to get/set data attributes,
 * not through workflow state transitions.
 */

const (
	WorkflowType = "s3-get-set-data-attributes"
	State1       = "S1"

	SmallDataKey        = "small-data"
	LargeDataKey        = "large-data"
	AnotherLargeDataKey = "another-large-data"

	// Small data content (stays in Temporal - under 50 byte threshold)
	SmallDataContent = "small"

	// Large data content (goes to S3 - over 50 byte threshold)
	LargeDataContent        = "large-data-content-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // Over 50 bytes
	AnotherLargeDataContent = "another-large-data-content-yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"         // Over 50 bytes

	// Updated values for testing updates
	UpdatedSmallDataContent = "updated-small"
	UpdatedLargeDataContent = "updated-large-data-content-zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz" // Over 50 bytes
)

var (
	SmallDataValue = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString("\"" + SmallDataContent + "\""),
	}

	LargeDataValue = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString("\"" + LargeDataContent + "\""),
	}

	AnotherLargeDataValue = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString("\"" + AnotherLargeDataContent + "\""),
	}

	UpdatedSmallDataValue = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString("\"" + UpdatedSmallDataContent + "\""),
	}

	UpdatedLargeDataValue = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString("\"" + UpdatedLargeDataContent + "\""),
	}
)

type handler struct {
	invokeHistory sync.Map
}

func NewHandler() *handler {
	_ = "STUB: not implemented"

	// GetTestResult returns the test result
	return nil
}

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApiV1WorkflowStartPost - Define workflow states
func (h *handler) ApiV1WorkflowStartPost(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// ApiV1WorkflowStateStart - Handle state start (waitUntil)
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Simple waitUntil - no commands, just proceed

// ApiV1WorkflowStateDecide - Handle state execution (execute)
func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Complete the workflow
