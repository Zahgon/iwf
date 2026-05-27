package integ

import (
	"net/http"
	"testing"

	"github.com/indeedeng/iwf/service/common/blobstore"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/common"
	"github.com/indeedeng/iwf/service"
	uclient "github.com/indeedeng/iwf/service/client"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
)

const testNamespace = "default"

func createTemporalClient(dataConverter converter.DataConverter) client.Client {
	_ = "STUB: not implemented"
	return *new(client.Client)
}

func startWorkflowWorkerWithRpc(handler common.WorkflowHandlerWithRpc, t *testing.T) (closeFunc func()) {
	_ = "STUB: not implemented"
	return nil
}

func startWorkflowWorker(handler common.WorkflowHandler, t *testing.T) (closeFunc func()) {
	_ = "STUB: not implemented"
	return nil
}

func doStartWorkflowWorker(handler common.WorkflowHandler, t *testing.T, router *gin.Engine) (closeFunc func()) {
	_ = "STUB: not implemented"
	return nil
}

type IwfServiceTestConfig struct {
	BackendType                      service.BackendType
	MemoEncryption                   bool
	DisableFailAtMemoIncompatibility bool // default to false so that we will fail at test
	DefaultHeaders                   map[string]string
	S3TestThreshold                  int
}

func startIwfService(backendType service.BackendType) (closeFunc func()) {
	_ = "STUB: not implemented"
	return nil
}

func startIwfServiceByConfig(config IwfServiceTestConfig) (uclient uclient.UnifiedClient, closeFunc func()) {
	_ = "STUB: not implemented"
	return *new(uclient.UnifiedClient), nil
}

func startIwfServiceWithClient(backendType service.BackendType) (uclient uclient.UnifiedClient, closeFunc func()) {
	_ = "STUB: not implemented"
	return *new(uclient.UnifiedClient), nil
}

//if backendType == service.BackendTypeTemporal {
//if integTemporalUclientCached == nil {
//	return doStartIwfServiceWithClient(backendType)
//}
//return integTemporalUclientCached, func() {}
//}
//if integCadenceUclientCached == nil {
//	return doStartIwfServiceWithClient(backendType)
//}
//return integCadenceUclientCached, func() {}

// disable caching for now as it makes it difficult to test memo
//var integCadenceUclientCached api.UnifiedClient
//var integTemporalUclientCached api.UnifiedClient

// globalBlobStore is a global var in this package for testing
var globalBlobStore blobstore.BlobStore

func doStartIwfServiceWithClient(config IwfServiceTestConfig) (uclient uclient.UnifiedClient, closeFunc func()) {
	_ = "STUB: not implemented"
	return *new(uclient.UnifiedClient), nil
}

// start iwf interpreter worker

// start iwf interpreter worker

func failTestAtError(err error, t *testing.T) { _ = "STUB: not implemented"; return }

func failTestAtHttpError(err error, httpResp *http.Response, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func failTestAtHttpErrorOrWorkflowUncompleted(err error, httpResp *http.Response, resp *iwfidl.WorkflowGetResponse, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func smallWaitForFastTest() { _ = "STUB: not implemented"; return }

func minimumContinueAsNewConfig(optimizeActivity bool) *iwfidl.WorkflowConfig {
	_ = "STUB: not implemented"
	return nil
}

func minimumGreedyTimerConfig() *iwfidl.WorkflowConfig { _ = "STUB: not implemented"; return nil }

func greedyTimerConfig(continueAsNew bool) *iwfidl.WorkflowConfig {
	_ = "STUB: not implemented"
	return nil
}

func minimumContinueAsNewConfigV0() *iwfidl.WorkflowConfig { _ = "STUB: not implemented"; return nil }

func getBackendTypes() []service.BackendType { _ = "STUB: not implemented"; return nil }
