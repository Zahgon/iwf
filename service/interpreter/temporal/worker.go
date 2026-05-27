package temporal

import (
	"github.com/indeedeng/iwf/service/common/blobstore"

	"github.com/indeedeng/iwf/config"
	uclient "github.com/indeedeng/iwf/service/client"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/worker"
)

type InterpreterWorker struct {
	temporalClient client.Client
	worker         worker.Worker
	taskQueue      string
}

func NewInterpreterWorker(
	config config.Config, temporalClient client.Client, taskQueue string, memoEncryption bool,
	memoEncryptionConverter converter.DataConverter, unifiedClient uclient.UnifiedClient,
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

// TODO: this cannot be too small otherwise the persistence_test for continueAsNew will fail, probably a bug in Temporal goSDK.
// It seems work as "parallelism" of something... need to report a bug ticket...

// When DisableStickyCache is true it can harm performance; should not be used in production environment

// TODO: remove in next release
// TODO: remove in next release
