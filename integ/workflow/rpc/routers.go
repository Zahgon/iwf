package rpc

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
 *		- WaitUntil updates attribute data and data objects and then waits until the channel has been published to
 * 		- Execute method moves to State2
 * State2:
 *		- WaitUntil method does nothing
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType              = "rpc"
	State1                    = "S1"
	State2                    = "S2"
	TestInterStateChannelName = "test-TestInterStateChannelName"
	RPCName                   = "test-RPCName"
	RPCNameReadOnly           = "test-RPC-readonly"
	RPCNameError              = "test-RPC-error"

	TestDataAttributeKey = "test-data-attribute"

	TestSearchAttributeKeywordKey    = "CustomKeywordField"
	TestSearchAttributeKeywordValue1 = "keyword-value1"
	TestSearchAttributeKeywordValue2 = "keyword-value2"

	TestSearchAttributeIntKey    = "CustomIntField"
	TestSearchAttributeBoolKey   = "CustomBoolField"
	TestSearchAttributeIntValue1 = 1
	TestSearchAttributeIntValue2 = 2

	WorkerApiErrorDetails = "test-details"
	WorkerApiErrorType    = "test-type"
)

type handler struct {
	invokeHistory sync.Map
	invokeData    sync.Map
}

func NewHandler() common.WorkflowHandlerWithRpc {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandlerWithRpc)
}

var TestDataAttributeVal1 = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-data-attribute-value1"),
}

var TestDataAttributeVal2 = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-data-attribute-value2"),
}

var TestInput = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-input-value"),
}

var TestOutput = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-output-value"),
}

var TestRecordEvent = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-record-event-value"),
}

var TestInterstateChannelValue = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-interstatechannel-value"),
}

func (h *handler) ApiV1WorkflowWorkerRpc(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Proceed with State 2 after setting the attributes

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Proceed after attributes and data objects have been updated and channel has been published to

// Go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to state 2

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
