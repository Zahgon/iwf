package rpcStorage

import (
	"strings"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/**
 * Test workflow for RPC external storage functionality.
 * Tests updating data attributes with both small and large data via RPC methods.
 *
 * State1:
 *   - Sets up initial data attributes (small and large)
 *   - Waits for RPC to update the data attributes
 * State2:
 *   - Completes the workflow
 */

const (
	WorkflowType            = "rpc-external-storage"
	State1                  = "S1"
	State2                  = "S2"
	UpdateDataAttributesRPC = "update-data-attributes"

	SmallDataKey = "small-data"
	LargeDataKey = "large-data"

	// Small data stays in Temporal (under threshold)
	SmallDataContent = "small-data-content"

	// Initial data for testing
	InitialSmallDataContent = "initial-small-data"
)

var (
	// Large data goes to external storage (over threshold) - 1KB+
	LargeDataContent = "large-data-content-" + strings.Repeat("x", 1000)

	// Initial large data for testing - 1KB+
	InitialLargeDataContent = "initial-large-data-" + strings.Repeat("y", 1000)
)

var (
	SmallDataValue = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString(SmallDataContent),
	}

	LargeDataValue = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString(LargeDataContent),
	}

	InitialSmallData = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString(InitialSmallDataContent),
	}

	InitialLargeData = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString(InitialLargeDataContent),
	}

	TestInput = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString("test-input-value"),
	}

	TestOutput = iwfidl.EncodedObject{
		Encoding: iwfidl.PtrString("json"),
		Data:     iwfidl.PtrString("test-output-value"),
	}
)

type handler struct {
	testData sync.Map
}

func NewHandler() common.WorkflowHandlerWithRpc {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandlerWithRpc)
}

func (h *handler) ApiV1WorkflowWorkerRpc(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Store received data for verification

// Verify we received the current data attributes (loaded from external storage)

// Verify we received actual data content, not just external storage references

// Update data attributes with new values and send signal to close workflow

func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Set up initial data attributes and wait for internal signal

// Final state - no commands needed

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Only complete workflow when we receive the close-workflow signal

// We received the internal signal to close workflow

// Should not happen - wait until signal is received

// Complete the workflow

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	// Return empty history (not tracking state invocations for this test)
	return nil, nil
}

// Return test data collected from RPC calls
