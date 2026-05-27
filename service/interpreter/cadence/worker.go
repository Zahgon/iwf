package cadence

import (
	"github.com/indeedeng/iwf/config"
	"github.com/indeedeng/iwf/service/common/blobstore"

	uclient "github.com/indeedeng/iwf/service/client"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/worker"
)

type InterpreterWorker struct {
	service   workflowserviceclient.Interface
	closeFunc func()
	domain    string
	worker    worker.Worker
	tasklist  string
}

func NewInterpreterWorker(
	config config.Config, service workflowserviceclient.Interface, domain, tasklist string, closeFunc func(),
	unifiedClient uclient.UnifiedClient,
	store blobstore.BlobStore,
) *InterpreterWorker {
	_ = "STUB: not implemented"
	return nil
}

func (iw *InterpreterWorker) Close() { _ = "STUB: not implemented"; return }

func (iw *InterpreterWorker) StartWithStickyCacheDisabledForTest() {
	_ = "STUB: not implemented"
	return
}

func (iw *InterpreterWorker) Start() { _ = "STUB: not implemented"; return }

func (iw *InterpreterWorker) start(disableStickyCache bool) { _ = "STUB: not implemented"; return }

// override default

// override default

// When DisableStickyCache is true it can harm performance; should not be used in production environment

// TODO: remove in next release
// TODO: remove in next release
