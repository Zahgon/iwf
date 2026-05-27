package locking

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
 *		- WaitUntil method does nothing
 * 		- Execute method will move to State Waiting, and 10 instances of State 2
 * State2:
 * 		- WaitUntil update SA
 * 		- Execute method will update data attributes and will gracefully complete workflow
 * StateWaiting:
 * 		- WaitUntil will proceed once the internal channel has been published to
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType                  = "locking"
	State1                        = "S1"
	State2                        = "S2"
	StateWaiting                  = "StateWaiting"
	TestDataAttributeKey1         = "test-data-attribute-1"
	TestDataAttributeKey2         = "test-data-attribute-2"
	RPCName                       = "increase-counter"
	InternalChannelName           = "test-channel"
	TestSearchAttributeKeywordKey = "CustomKeywordField"
	TestSearchAttributeIntKey     = "CustomIntField"

	ShouldUnblockStateWaiting = "shouldUnblockStateWaiting"

	InParallelS2 = 10

	NumUnusedSignals = 4

	UnusedSignalChannelName   = "test-unused-signal-channel"
	UnusedInternalChannelName = "test-unused-internal-channel"
)

var TestValue = &iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("data"),
}

var UnblockValue = &iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString(ShouldUnblockStateWaiting),
}

var state2Options = &iwfidl.WorkflowStateOptions{
	SearchAttributesLoadingPolicy: &iwfidl.PersistenceLoadingPolicy{
		PersistenceLoadingType: iwfidl.PARTIAL_WITH_EXCLUSIVE_LOCK.Ptr(),
		PartialLoadingKeys: []string{
			TestSearchAttributeIntKey,
			TestSearchAttributeKeywordKey,
		},
		LockingKeys: []string{
			TestSearchAttributeIntKey,
		},
	},
	DataAttributesLoadingPolicy: &iwfidl.PersistenceLoadingPolicy{
		PersistenceLoadingType: iwfidl.PARTIAL_WITH_EXCLUSIVE_LOCK.Ptr(),
		PartialLoadingKeys: []string{
			TestDataAttributeKey1,
			TestDataAttributeKey2,
		},
		LockingKeys: []string{
			TestDataAttributeKey1,
		},
	},
}

var state2Movement = iwfidl.StateMovement{
	StateId:      State2,
	StateOptions: state2Options,
}

type handler struct {
	invokeHistory sync.Map
	rpcInvokes    int32
}

func NewHandler() common.WorkflowHandlerWithRpc {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandlerWithRpc)
}

func (h *handler) ApiV1WorkflowWorkerRpc(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Publish to internal channel

// the 4 messages are sent from the beginning of "locking_test"

// This RPC will increase both SA and DA

// ApiV1WorkflowStateStart - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Go straight to the decide methods without any commands

// Will proceed once the internal channel has been published to

// This state API is to increase SA

// Go straight to the decide methods after updating SA

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Move to State Waiting, and 10 instances of State 2
// State Waiting will not complete until the internal channel has been published to

// Move to completion

// This API is to increase DA

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
