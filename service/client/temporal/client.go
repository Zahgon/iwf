package temporal

import (
	"context"

	"github.com/indeedeng/iwf/config"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	uclient "github.com/indeedeng/iwf/service/client"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
)

type temporalClient struct {
	tClient                        client.Client
	namespace                      string
	dataConverter                  converter.DataConverter
	memoEncryption                 bool // this is a workaround for https://github.com/temporalio/sdk-go/issues/1045
	queryWorkflowFailedRetryPolicy config.QueryWorkflowFailedRetryPolicy
}

func NewTemporalClient(
	tClient client.Client, namespace string, dataConverter converter.DataConverter, memoEncryption bool, retryPolicy *config.QueryWorkflowFailedRetryPolicy,
) uclient.UnifiedClient {
	_ = "STUB: not implemented"
	return *new(uclient.UnifiedClient)
}

func (t *temporalClient) Close() { _ = "STUB: not implemented"; return }

func (t *temporalClient) IsWorkflowAlreadyStartedError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

// there is no type to check, just a string
// https://github.com/temporalio/sdk-go/blob/d10e87118a07b44fd09bf88d39a628f0e6e70c34/internal/error.go#L336

func (t *temporalClient) GetRunIdFromWorkflowAlreadyStartedError(err error) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (t *temporalClient) IsNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

func (t *temporalClient) isQueryFailedError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *temporalClient) IsRequestTimeoutError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *temporalClient) IsWorkflowTimeoutError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *temporalClient) GetApplicationErrorTypeIfIsApplicationError(err error) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *temporalClient) GetApplicationErrorDetails(err error, detailsPtr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *temporalClient) GetApplicationErrorTypeAndDetails(err error) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// All other types, e.g. iwfidl.StateCompletionOutput, try to Marshal the object to JSON

func (t *temporalClient) StartInterpreterWorkflow(
	ctx context.Context, options uclient.StartWorkflowOptions, args ...interface{},
) (runId string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// use temporal schedule instead of cron
// https://temporal.io/blog/how-do-i-convert-my-cron-into-a-schedule
// workflowOptions.CronSchedule = *options.CronSchedule

func (t *temporalClient) StartWaitForStateCompletionWorkflow(
	ctx context.Context, options uclient.StartWorkflowOptions,
) (runId string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// the workflow could be timeout, so we allow duplicate

// because of WorkflowExecutionErrorWhenAlreadyStarted: false, we won't get WorkflowAlreadyStartedError as we do in Cadence

func (t *temporalClient) StartBlobStoreCleanupWorkflow(
	ctx context.Context, taskQueue, workflowID, cronSchedule, storeId string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *temporalClient) SignalWithStartWaitForStateCompletionWorkflow(
	ctx context.Context, options uclient.StartWorkflowOptions, stateCompletionOutput iwfidl.StateCompletionOutput,
) error {
	_ = "STUB: not implemented"
	return nil
}

// the workflow could be timeout, so we allow duplicate

func (t *temporalClient) SignalWorkflow(
	ctx context.Context, workflowID string, runID string, signalName string, arg interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *temporalClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *temporalClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *temporalClient) ListWorkflow(
	ctx context.Context, request *uclient.ListWorkflowExecutionsRequest,
) (*uclient.ListWorkflowExecutionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *temporalClient) QueryWorkflow(
	ctx context.Context, valuePtr interface{}, workflowID string, runID string, queryType string, args ...interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

// Only QueryFailed error causes retry; all other errors make the loop to finish immediately

func (t *temporalClient) DescribeWorkflowExecution(
	ctx context.Context, workflowID, runID string, requestedSearchAttributes []iwfidl.SearchAttributeKeyAndType,
) (*uclient.DescribeWorkflowExecutionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *temporalClient) encryptMemoIfNeeded(rawMemo map[string]interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *temporalClient) getMemoAndDecryptIfNeeded(memo *common.Memo) (map[string]iwfidl.EncodedObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mapToTemporalWorkflowIdReusePolicy(workflowIdReusePolicy iwfidl.WorkflowIDReusePolicy) (*enums.WorkflowIdReusePolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mapToIwfWorkflowStatus(status enums.WorkflowExecutionStatus) (iwfidl.WorkflowStatus, error) {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowStatus), nil
}

func (t *temporalClient) GetWorkflowResult(
	ctx context.Context, valuePtr interface{}, workflowID string, runID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *temporalClient) SynchronousUpdateWorkflow(
	ctx context.Context, valuePtr interface{}, workflowID, runID, updateType string, input interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Leaving this as Accepted that was a default value before WaitForStage became required argument, but Completed might be a better choice

func (t *temporalClient) ResetWorkflow(
	ctx context.Context, request iwfidl.WorkflowResetRequest,
) (runId string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// set default runId to current

func (t *temporalClient) GetBackendType() (backendType service.BackendType) {
	_ = "STUB: not implemented"
	return *new(service.BackendType)
}

func (t *temporalClient) GetApiService() interface{} { _ = "STUB: not implemented"; return nil }
