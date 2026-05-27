package compatibility

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
)

func GetStartApiTimeoutSeconds(stateOptions *iwfidl.WorkflowStateOptions) int32 {
	_ = "STUB: not implemented"
	return 0
}

func GetDecideApiTimeoutSeconds(stateOptions *iwfidl.WorkflowStateOptions) int32 {
	_ = "STUB: not implemented"
	return 0
}

func GetStartApiRetryPolicy(stateOptions *iwfidl.WorkflowStateOptions) *iwfidl.RetryPolicy {
	_ = "STUB: not implemented"
	return nil
}

func GetDecideApiRetryPolicy(stateOptions *iwfidl.WorkflowStateOptions) *iwfidl.RetryPolicy {
	_ = "STUB: not implemented"
	return nil
}

func GetWaitUntilApiDataAttributesLoadingPolicy(stateOptions *iwfidl.WorkflowStateOptions) *iwfidl.PersistenceLoadingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func GetExecuteApiDataAttributesLoadingPolicy(stateOptions *iwfidl.WorkflowStateOptions) *iwfidl.PersistenceLoadingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func GetWaitUntilApiSearchAttributesLoadingPolicy(stateOptions *iwfidl.WorkflowStateOptions) *iwfidl.PersistenceLoadingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func GetExecuteApiSearchAttributesLoadingPolicy(stateOptions *iwfidl.WorkflowStateOptions) *iwfidl.PersistenceLoadingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func GetStartApiFailurePolicy(stateOptions *iwfidl.WorkflowStateOptions) *iwfidl.StartApiFailurePolicy {
	_ = "STUB: not implemented"
	return nil
}

func GetSkipWaitUntilApi(stateOptions *iwfidl.WorkflowStateOptions) bool {
	_ = "STUB: not implemented"
	return false
}
