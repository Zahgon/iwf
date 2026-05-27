package retry

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
	"go.temporal.io/sdk/temporal"
	"go.uber.org/cadence/workflow"
)

func ConvertCadenceWorkflowRetryPolicy(policy *iwfidl.WorkflowRetryPolicy) *workflow.RetryPolicy {
	_ = "STUB: not implemented"
	return nil
}

func ConvertCadenceActivityRetryPolicy(policy *iwfidl.RetryPolicy) *workflow.RetryPolicy {
	_ = "STUB: not implemented"
	return nil
}

// in Cadence, ExpirationInterval is the timeout include all retries

// unlimited to match Temporal

func ConvertTemporalWorkflowRetryPolicy(policy *iwfidl.WorkflowRetryPolicy) *temporal.RetryPolicy {
	_ = "STUB: not implemented"
	return nil
}

func ConvertTemporalActivityRetryPolicy(policy *iwfidl.RetryPolicy) *temporal.RetryPolicy {
	_ = "STUB: not implemented"
	return nil
}

func fillActivityRetryPolicyDefault(policy *iwfidl.RetryPolicy) { _ = "STUB: not implemented"; return }

func fillWorkflowRetryPolicyDefault(policy *iwfidl.WorkflowRetryPolicy) {
	_ = "STUB: not implemented"
	return
}
