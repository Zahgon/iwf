package compatibility

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
)

func GetWorkflowIdReusePolicy(options iwfidl.WorkflowStartOptions) *iwfidl.WorkflowIDReusePolicy {
	_ = "STUB: not implemented"
	return nil
}

// Keeping typo enum for backwards compatibility. Both old and corrected enums return the same result.
