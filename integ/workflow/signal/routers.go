package signal

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/common"
	"github.com/indeedeng/iwf/service/common/ptr"
)

/**
 * This test workflow has two states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil waits until 4 signals are received
 * 		- Execute method publishes the 4 signals & moves to State2
 * State2:
 *		- WaitUntil method does nothing
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType                  = "signal"
	State1                        = "S1"
	State2                        = "S2"
	SignalName                    = "test-signal-name"
	InternalChannelName           = "test-internal-channel-name"
	UnhandledSignalName           = "test-unhandled-signal-name"
	RPCNameGetSignalChannelInfo   = "RPCNameGetSignalChannelInfo"
	RPCNameGetInternalChannelInfo = "RPCNameGetInternalChannelInfo"
)

var StateOptionsForLargeDataAttributes = iwfidl.WorkflowStateOptions{
	DataAttributesLoadingPolicy: &iwfidl.PersistenceLoadingPolicy{
		PersistenceLoadingType: ptr.Any(iwfidl.NONE),
	},
}

type handler struct {
	invokeHistory sync.Map
	invokeData    sync.Map
}

func NewHandler() common.WorkflowHandlerWithRpc {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandlerWithRpc)
}

func (h *handler) ApiV1WorkflowWorkerRpc(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Proceed when 4 signals are received

// Go straight to the decide methods without any commands

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Publish 4 signals

// Move to State 2

// Move to completion

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
