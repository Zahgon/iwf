package interstate

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
)

/**
 * This test workflow has four states, using REST controller to implement the workflow directly.
 *
 * State1:
 *		- WaitUntil method does nothing
 * 		- Execute method will move to State21 & State22:
 * State21:
 * 		- WaitUntil will proceed once channel1 has been published to
 * 		- Execute method will move to State31:
 * State22:
 * 		- WaitUntil will delay 2s then publish on channel1
 *      - Execute method will delay 2s then publish on channel2 & end in a dead-end
 * State31:
 * 		- WaitUntil will proceed once channel2 has been published to
 *      - Execute method will gracefully complete workflow
 */
const (
	WorkflowType = "interstate"
	State1       = "S1"
	State21      = "S21"
	State22      = "S22"
	State31      = "S31"

	channel1 = "channel1"
	channel2 = "channel2"
)

var TestVal1 = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-value1"),
}

var TestVal2 = iwfidl.EncodedObject{
	Encoding: iwfidl.PtrString("json"),
	Data:     iwfidl.PtrString("test-value2"),
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

// Go straight to the decide methods without any commands

// Will proceed once channel 1 has been published to

// Will proceed once channel 2 has been published to

// Wait 2 seconds then publish on channel1

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// State 1 requires no pre-reqs
// Move to state 21 & 22:
// 21 - Will wait for channel 1
// 22 - Will wait 3 seconds then publish to channel 1

// Move to state 31, which will wait for channel 2

// Move to completion

// Move to the dead-end state and publish on channel 2 (to unlock State 31)

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
