package conditional_close

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * This test workflow has 1 state, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil will proceed when the channel or signal is published to
 *      - Execute method will continuously retry State1 until the 3rd attempt which will send a message to the channel or
 *        signal, making the state empty and force-complete.
 */
const (
	WorkflowType              = "conditional_close"
	RpcPublishInternalChannel = "publish_internal_channel"

	TestChannelName = "test-channel-name"

	State1 = "S1"
)

var TestInput = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-data"),
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

// Return channel name

// ApiV1WorkflowStateStart - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Proceed when channel is published to

// Proceed when signal is published to

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Wait for 3 seconds so that the channel can have a new message

// Send internal channel message within the state execution
// and expecting the messages are processed by the conditional check

// Use signal instead

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
