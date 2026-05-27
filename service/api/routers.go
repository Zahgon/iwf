package api

import (
	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/config"
	uclient "github.com/indeedeng/iwf/service/client"
	"github.com/indeedeng/iwf/service/common/blobstore"
	"github.com/indeedeng/iwf/service/common/log"
)

const WorkflowStartApiPath = "/api/v1/workflow/start"
const WorkflowWaitForStateCompletionApiPath = "/api/v1/workflow/waitForStateCompletion"
const WorkflowSignalApiPath = "/api/v1/workflow/signal"
const WorkflowPublishToInternalChannelApiPath = "/api/v1/workflow/publishToInternalChannel"
const WorkflowGetDataAttributesApiPath = "/api/v1/workflow/dataobjects/get"
const WorkflowSetDataAttributesApiPath = "/api/v1/workflow/dataobjects/set"
const WorkflowGetSearchAttributesApiPath = "/api/v1/workflow/searchattributes/get"
const WorkflowSetSearchAttributesApiPath = "/api/v1/workflow/searchattributes/set"
const WorkflowGetApiPath = "/api/v1/workflow/get"
const WorkflowGetWithWaitApiPath = "/api/v1/workflow/getWithWait"
const WorkflowSearchApiPath = "/api/v1/workflow/search"
const WorkflowResetApiPath = "/api/v1/workflow/reset"
const WorkflowSkipTimerApiPath = "/api/v1/workflow/timer/skip"
const WorkflowStopApiPath = "/api/v1/workflow/stop"
const WorkflowInternalDumpApiPath = "/api/v1/workflow/internal/dump"
const WorkflowConfigUpdateApiPath = "/api/v1/workflow/config/update"
const WorkflowTriggerContinueAsNewApiPath = "/api/v1/workflow/triggerContinueAsNew"
const WorkflowRpcApiPath = "/api/v1/workflow/rpc"
const InfoHealthCheck = "/info/healthcheck"

// NewService returns a new router.
func NewService(config config.Config, client uclient.UnifiedClient, logger log.Logger, store blobstore.BlobStore) *gin.Engine {
	_ = "STUB: not implemented"
	return nil
}
