package cadence

import (
	"context"

	"github.com/indeedeng/iwf/config"

	"github.com/indeedeng/iwf/service"

	"github.com/indeedeng/iwf/gen/iwfidl"
	uclient "github.com/indeedeng/iwf/service/client"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/encoded"
)

type cadenceClient struct {
	domain                         string
	cClient                        client.Client
	closeFunc                      func()
	serviceClient                  workflowserviceclient.Interface
	converter                      encoded.DataConverter
	queryWorkflowFailedRetryPolicy config.QueryWorkflowFailedRetryPolicy
}

func (t *cadenceClient) IsWorkflowAlreadyStartedError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *cadenceClient) GetRunIdFromWorkflowAlreadyStartedError(err error) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (t *cadenceClient) IsNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

func (t *cadenceClient) isQueryFailedError(err error) bool { _ = "STUB: not implemented"; return false }

func (t *cadenceClient) IsWorkflowTimeoutError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *cadenceClient) IsRequestTimeoutError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *cadenceClient) GetApplicationErrorTypeIfIsApplicationError(err error) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *cadenceClient) GetApplicationErrorDetails(err error, detailsPtr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *cadenceClient) GetApplicationErrorTypeAndDetails(err error) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// All other types, e.g. iwfidl.StateCompletionOutput, try to Marshal the object to JSON

func NewCadenceClient(
	domain string, cClient client.Client, serviceClient workflowserviceclient.Interface,
	converter encoded.DataConverter, closeFunc func(), retryPolicy *config.QueryWorkflowFailedRetryPolicy,
) uclient.UnifiedClient {
	_ = "STUB: not implemented"
	return *new(uclient.UnifiedClient)
}

func (t *cadenceClient) Close() { _ = "STUB: not implemented"; return }

func (t *cadenceClient) StartInterpreterWorkflow(
	ctx context.Context, options uclient.StartWorkflowOptions, args ...interface{},
) (runId string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *cadenceClient) StartWaitForStateCompletionWorkflow(
	ctx context.Context, options uclient.StartWorkflowOptions,
) (runId string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// the workflow could be timeout, so we allow duplicate

// if the workflow is already started, we return the runId

func (t *cadenceClient) StartBlobStoreCleanupWorkflow(
	ctx context.Context, taskQueue, workflowID, cronSchedule, storeId string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *cadenceClient) SignalWithStartWaitForStateCompletionWorkflow(
	ctx context.Context, options uclient.StartWorkflowOptions, stateCompletionOutput iwfidl.StateCompletionOutput,
) error {
	_ = "STUB: not implemented"
	return nil
}

// the workflow could be timeout, so we allow duplicate

func (t *cadenceClient) SignalWorkflow(
	ctx context.Context, workflowID string, runID string, signalName string, arg interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *cadenceClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *cadenceClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *cadenceClient) ListWorkflow(
	ctx context.Context, request *uclient.ListWorkflowExecutionsRequest,
) (*uclient.ListWorkflowExecutionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *cadenceClient) QueryWorkflow(
	ctx context.Context, valuePtr interface{}, workflowID string, runID string, queryType string, args ...interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

// Only QueryFailed error causes retry; all other errors make the loop to finish immediately

func queryWorkflowWithStrongConsistency(
	t *cadenceClient, ctx context.Context, workflowID string, runID string, queryType string, args []interface{},
) (encoded.Value, error) {
	_ = "STUB: not implemented"
	return *new(encoded.Value), nil
}

func (t *cadenceClient) DescribeWorkflowExecution(
	ctx context.Context, workflowID, runID string, requestedSearchAttributes []iwfidl.SearchAttributeKeyAndType,
) (*uclient.DescribeWorkflowExecutionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cadence does not provide FirstRunId

func (t *cadenceClient) decodeMemo(memo *shared.Memo) (map[string]iwfidl.EncodedObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mapToCadenceWorkflowIdReusePolicy(workflowIdReusePolicy iwfidl.WorkflowIDReusePolicy) (*client.WorkflowIDReusePolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mapToIwfWorkflowStatus(status *shared.WorkflowExecutionCloseStatus) (iwfidl.WorkflowStatus, error) {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowStatus), nil
}

func (t *cadenceClient) GetWorkflowResult(
	ctx context.Context, valuePtr interface{}, workflowID string, runID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *cadenceClient) SynchronousUpdateWorkflow(
	ctx context.Context, valuePtr interface{}, workflowID, runID, updateType string, input interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *cadenceClient) ResetWorkflow(
	ctx context.Context, request iwfidl.WorkflowResetRequest,
) (newRunId string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// set default runId to current

// TODO not sure why Cadence reset API requires this for GetWorkflowExecutionHistory API....

func (t *cadenceClient) GetBackendType() (backendType service.BackendType) {
	_ = "STUB: not implemented"
	return *new(service.BackendType)
}

func (t *cadenceClient) GetApiService() interface{} { _ = "STUB: not implemented"; return nil }
