package api

import (
	"github.com/indeedeng/iwf/config"
	"github.com/indeedeng/iwf/service/common/blobstore"

	uclient "github.com/indeedeng/iwf/service/client"
	"github.com/indeedeng/iwf/service/common/errors"
	"github.com/indeedeng/iwf/service/common/log"

	"github.com/gin-gonic/gin"
)

type handler struct {
	svc    ApiService
	logger log.Logger
}

func newHandler(config config.Config, client uclient.UnifiedClient, logger log.Logger, store blobstore.BlobStore) *handler {
	_ = "STUB: not implemented"
	return nil
}

func (h *handler) close() {
	_ = "STUB: not implemented"

	// Index is the index handler.
	return
}

func (h *handler) index(c *gin.Context) { _ = "STUB: not implemented"; return }

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) apiV1WorkflowStart(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowWaitForStateCompletion(c *gin.Context) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) apiV1WorkflowSignal(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1PublishToInternalChannel(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowStop(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowInternalDump(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowConfigUpdate(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowTriggerContinueAsNew(c *gin.Context) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) apiV1WorkflowSearch(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowRpc(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) infoHealthCheck(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowGetDataAttributes(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowSetDataAttributes(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowGetSearchAttributes(c *gin.Context) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) apiV1WorkflowSetSearchAttributes(c *gin.Context) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) apiV1WorkflowGet(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowGetWithWait(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) doApiV1WorkflowGetPost(c *gin.Context, waitIfStillRunning bool) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) apiV1WorkflowReset(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) apiV1WorkflowSkipTimer(c *gin.Context) { _ = "STUB: not implemented"; return }

func invalidRequestSchema(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) processError(c *gin.Context, resp *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return
}
