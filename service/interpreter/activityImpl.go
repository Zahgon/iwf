package interpreter

import (
	"context"
	"net/http"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

// StateStart is Deprecated, will be removed in next release
func StateStart(
	ctx context.Context, backendType service.BackendType, input service.StateStartActivityInput, searchAttributes []iwfidl.SearchAttribute,
) (*iwfidl.WorkflowStateStartResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func StateApiWaitUntil(
	ctx context.Context, backendType service.BackendType, input service.StateStartActivityInput, searchAttributes []iwfidl.SearchAttribute,
) (*iwfidl.WorkflowStateStartResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load data attributes from external storage

// Before returning successful results, check if it's local activity then compose some info for debug purpose
// This is because local activity doesn't record input into the history.
// But there are some small info that are important to record

// StateDecide is deprecated. Will be removed in next release
func StateDecide(
	ctx context.Context,
	backendType service.BackendType,
	input service.StateDecideActivityInput,
) (*iwfidl.WorkflowStateDecideResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func StateApiExecute(
	ctx context.Context,
	backendType service.BackendType,
	input service.StateDecideActivityInput,
) (*iwfidl.WorkflowStateDecideResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load data attributes from external storage

// workflowImpl is only passing StateWaitUntilFailed, not StateStartApiSucceeded, to save the history from storing duplicate data
// So we need to construct the other for backward compatibility(to old SDK that is still using StateStartApiSucceeded)

// Before returning successful results, check if it's local activity then compose some info for debug purpose
// This is because local activity doesn't record input into the history.
// But there are some small info that are important to record

// Externalize only when enabled and blob store is available (nil when e.g. STAGING_LEVEL was empty at worker start).

func composeInputForDebug(stateExeId string) *string {
	_ = "STUB: not implemented"
	// NOTE: only use the stateExecutionId for now, but we can add more later if needed
	return nil
}

func checkStateDecisionFromResponse(resp *iwfidl.WorkflowStateDecideResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func printDebugMsg(logger interfaces.UnifiedLogger, err error, url string) {
	_ = "STUB: not implemented"
	return
}

func composeStartApiRespError(provider interfaces.ActivityProvider, err error, resp *iwfidl.WorkflowStateStartResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func composeExecuteApiRespError(provider interfaces.ActivityProvider, err error, resp *iwfidl.WorkflowStateDecideResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func checkHttpError(err error, httpResp *http.Response) bool {
	_ = "STUB: not implemented"
	return false
}

func composeHttpError(
	isLocalActivity bool, provider interfaces.ActivityProvider, err error, httpResp *http.Response, errType string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func trimText(msg string, maxLength int) string { _ = "STUB: not implemented"; return "" }

func checkCommandRequestFromWaitUntilResponse(resp *iwfidl.WorkflowStateStartResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// every command must have an id for this type

// Check if each command in the combinations has a matching command in one of the lists

// NOTE: we don't require decider trigger type when there is no commands

func areAllCommandCombinationsIdsValid(commandReq *iwfidl.CommandRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func listTimerSignalInternalChannelCommandIds(commandReq *iwfidl.CommandRequest) []string {
	_ = "STUB: not implemented"
	return nil
}

func DumpWorkflowInternal(
	ctx context.Context, backendType service.BackendType, req iwfidl.WorkflowDumpRequest,
) (*iwfidl.WorkflowDumpResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InvokeWorkerRpc(
	ctx context.Context, backendType service.BackendType, rpcPrep *service.PrepareRpcQueryResponse,
	req iwfidl.WorkflowRpcRequest,
) (*interfaces.InvokeRpcActivityOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeNextStateInputsToExternalStorage(ctx context.Context, nextStates []iwfidl.StateMovement, currentInputCopy *iwfidl.EncodedObject, workflowId string) ([]iwfidl.StateMovement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processNextStateInputForExternalStorage(ctx context.Context, nextStateInput *iwfidl.EncodedObject, currentInputCopy *iwfidl.EncodedObject, workflowId string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if external storage is needed

// Try to reuse existing external storage if data is identical

// Reuse existing external storage

// Save to external storage

// loadStateInputFromExternalStorage is specifically for loading state input from external storage.
// It loads the data and replace the field of input,
// and it also preserves the external storage identifiers for potential reuse optimization by returning a "whole" encodedObject
func loadStateInputFromExternalStorage(ctx context.Context, input *iwfidl.EncodedObject) (*iwfidl.EncodedObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load the data but preserve external storage identifiers for potential reuse
// This allows us to optimize by reusing the same external storage location
// if the next state input has identical data

// doLoadFromExternalStorage loads data from external storage and replace the fields of the input
func doLoadFromExternalStorage(ctx context.Context, input *iwfidl.EncodedObject) error {
	_ = "STUB: not implemented"
	return nil
}

func CleanupBlobStore(
	ctx context.Context, backendType service.BackendType, storeId string,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Check if workflow still exists in Temporal
// We must always check because workflows can run indefinitely,
// and the retention period only applies after workflow closure

// Workflow has been removed from Temporal, safe to delete S3 objects

// If no error, workflow still exists (open or within retention), don't delete

// this is a long running activity
// using record heartbeat so that it won't timeout at startToClose timeout
