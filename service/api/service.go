package api

import (
	"context"

	"github.com/indeedeng/iwf/service/common/blobstore"

	"github.com/indeedeng/iwf/config"

	uclient "github.com/indeedeng/iwf/service/client"
	"github.com/indeedeng/iwf/service/common/errors"
	"github.com/indeedeng/iwf/service/common/log"

	"github.com/indeedeng/iwf/gen/iwfidl"
)

type serviceImpl struct {
	client    uclient.UnifiedClient
	store     blobstore.BlobStore
	taskQueue string
	logger    log.Logger
	config    config.Config
}

func (s *serviceImpl) Close() { _ = "STUB: not implemented"; return }

func NewApiService(
	cfg config.Config, client uclient.UnifiedClient, taskQueue string, logger log.Logger, store blobstore.BlobStore,
) (ApiService, error) {
	_ = "STUB: not implemented"
	return *new(ApiService), nil
}

func (s *serviceImpl) ApiV1WorkflowStartPost(
	ctx context.Context, req iwfidl.WorkflowStartRequest,
) (wresp *iwfidl.WorkflowStartResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

// workerUrl is always needed, for optimizing None as persistence loading type

// inject some code to upload the large input to S3 and replace the input

// Note: the value is actually not too important, we will check the presence of the key only as today

// inject some code to upload the large input to S3 and replace the input

// this feature requires workflowID not contains certain characters

// 1. check the size of the input is larger than the threshold

// 2. if it is, upload the input to S3

// 3. replace the input with the S3 object

func overrideWorkflowConfig(configOverride iwfidl.WorkflowConfig, workflowConfig *iwfidl.WorkflowConfig) {
	_ = "STUB: not implemented"
	return
}

func (s *serviceImpl) ApiV1WorkflowWaitForStateCompletion(
	ctx context.Context, req iwfidl.WorkflowWaitForStateCompletionRequest,
) (wresp *iwfidl.WorkflowWaitForStateCompletionResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

// waitForOn == "new"

// Temporal

// Cadence

// TODO: https://github.com/indeedeng/iwf-java-sdk/issues/218
// it doesn't seem to have a way for SDK to know the timeout at this API
// So hardcoded to 1 hour for now. If it timeouts, the IDReusePolicy will restart a new one

// the workflow is still running, but the wait has exceeded limit

func (s *serviceImpl) ApiV1WorkflowSignalPost(
	ctx context.Context, req iwfidl.WorkflowSignalRequest,
) (retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil
}

func (s *serviceImpl) ApiV1WorkflowPublishToInternalChannelPost(
	ctx context.Context, req iwfidl.PublishToInternalChannelRequest,
) (retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil
}

func (s *serviceImpl) ApiV1WorkflowConfigUpdate(
	ctx context.Context, req iwfidl.WorkflowConfigUpdateRequest,
) (retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil
}

func (s *serviceImpl) ApiV1WorkflowTriggerContinueAsNew(
	ctx context.Context, req iwfidl.TriggerContinueAsNewRequest,
) (retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil
}

func (s *serviceImpl) ApiV1WorkflowStopPost(
	ctx context.Context, req iwfidl.WorkflowStopRequest,
) (retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil
}

func (s *serviceImpl) ApiV1WorkflowGetQueryAttributesPost(
	ctx context.Context, req iwfidl.WorkflowGetDataObjectsRequest,
) (wresp *iwfidl.WorkflowGetDataObjectsResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note that when the requested keys is empty, it means all

// using memo is enough

// this means that we cannot use memo to continue, need to fall back to use query

// Load data from external storage if necessary

func (s *serviceImpl) ApiV1WorkflowSetQueryAttributesPost(
	ctx context.Context, req iwfidl.WorkflowSetDataObjectsRequest,
) (retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil
}

// Process data objects for external storage if necessary

func (s *serviceImpl) ApiV1WorkflowGetSearchAttributesPost(
	ctx context.Context, req iwfidl.WorkflowGetSearchAttributesRequest,
) (wresp *iwfidl.WorkflowGetSearchAttributesResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serviceImpl) ApiV1WorkflowSetSearchAttributesPost(
	ctx context.Context, req iwfidl.WorkflowSetSearchAttributesRequest,
) (retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil
}

func (s *serviceImpl) ApiV1WorkflowGetPost(
	ctx context.Context, req iwfidl.WorkflowGetRequest,
) (wresp *iwfidl.WorkflowGetResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serviceImpl) ApiV1WorkflowGetWithWaitPost(
	ctx context.Context, req iwfidl.WorkflowGetRequest,
) (wresp *iwfidl.WorkflowGetResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

// withWait:
//
//	 because s.client.GetWorkflowResult will wait for the completion if workflow is running --
//		when withWait is false, if workflow is not running and needResults is true, it will then call s.client.GetWorkflowResult to get results
//		when withWait is true, it will do everything
func (s *serviceImpl) doApiV1WorkflowGetPost(
	ctx context.Context, req iwfidl.WorkflowGetRequest, withWait bool,
) (wresp *iwfidl.WorkflowGetResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the workflow is still running, but the wait has exceeded limit

// workflow failed by interpreter decision, or by user workflow state decision

// it could be timeout/terminated/canceled/etc. We need to describe again to get the final status

// we cannot return these status, which will be a wrong results
// TODO: maybe return 4xx

func (s *serviceImpl) ApiV1WorkflowSearchPost(
	ctx context.Context, req iwfidl.WorkflowSearchRequest,
) (wresp *iwfidl.WorkflowSearchResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serviceImpl) ApiV1WorkflowRpcPost(
	ctx context.Context, req iwfidl.WorkflowRpcRequest,
) (wresp *iwfidl.WorkflowRpcResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

// search attributes are not available at this time

// if there is no mutation on the workflow, this RPC is "readonly", then don't send the signal

// send the signal

// the input/output is only for debugging purpose but could be too expensive to store

func needLocking(req iwfidl.WorkflowRpcRequest) bool { _ = "STUB: not implemented"; return false }

func (s *serviceImpl) handleRpcBySynchronousUpdate(
	ctx context.Context, req iwfidl.WorkflowRpcRequest,
) (resp *iwfidl.WorkflowRpcResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

func doNeedLocking(policy *iwfidl.PersistenceLoadingPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *serviceImpl) ApiV1WorkflowResetPost(
	ctx context.Context, req iwfidl.WorkflowResetRequest,
) (wresp *iwfidl.WorkflowResetResponse, retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serviceImpl) ApiV1WorkflowSkipTimerPost(
	ctx context.Context, request iwfidl.WorkflowSkipTimerRequest,
) (retError *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil
}

func (s *serviceImpl) ApiV1WorkflowDumpPost(
	ctx context.Context, request iwfidl.WorkflowDumpRequest,
) (*iwfidl.WorkflowDumpResponse, *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serviceImpl) ApiInfoHealth(ctx context.Context) *iwfidl.HealthInfo {
	_ = "STUB: not implemented"
	return nil
}

func makeInvalidRequestError(msg string) *errors.ErrorAndStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *serviceImpl) handleError(err error, apiPath string, workflowId string) *errors.ErrorAndStatus {
	_ = "STUB: not implemented"
	return nil
}

// writeDataObjectsToExternalStorage processes data objects and writes large ones to external storage
func (s *serviceImpl) writeDataObjectsToExternalStorage(ctx context.Context, dataObjects []iwfidl.KeyValue, workflowId string) ([]iwfidl.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clear data since it's now in external storage

// loadDataObjectsFromExternalStorage loads data from external storage for data objects that have external storage references
func (s *serviceImpl) loadDataObjectsFromExternalStorage(ctx context.Context, dataObjects []iwfidl.KeyValue) ([]iwfidl.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Preserve external storage reference
// Preserve external storage reference
