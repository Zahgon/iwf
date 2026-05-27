package env

import (
	"github.com/indeedeng/iwf/config"
	uclient "github.com/indeedeng/iwf/service/client"
	"github.com/indeedeng/iwf/service/common/blobstore"
	"go.temporal.io/sdk/converter"
)

// this env package is for creating some handy global objects so that workflow worker can make use of,
// to avoid passing throw worker context(esp hard because we have to do it with both Cadence & Temporal)
// Note it's not a good practice to use global so should keep them minimum

var sharedConfig config.Config

var temporalDataConverter converter.DataConverter

var temporalMemoEncryption bool

var unifiedClient uclient.UnifiedClient

var taskQueue string

var blobStore blobstore.BlobStore

func SetSharedEnv(
	config config.Config,
	memoEncryption bool,
	temporalMemoEncryptionDataConverter converter.DataConverter,
	client uclient.UnifiedClient,
	queue string,
	store blobstore.BlobStore,
) {
	_ = "STUB: not implemented"
	return
}

func GetUnifiedClient() uclient.UnifiedClient {
	_ = "STUB: not implemented"
	return *new(uclient.UnifiedClient)
}

func GetTaskQueue() string { _ = "STUB: not implemented"; return "" }

func GetSharedConfig() config.Config { _ = "STUB: not implemented"; return *new(config.Config) }

func CheckAndGetTemporalMemoEncryptionDataConverter() (converter.DataConverter, bool) {
	_ = "STUB: not implemented"
	return *new(converter.DataConverter), false
}

func GetBlobStore() blobstore.BlobStore {
	_ = "STUB: not implemented"
	return *new(blobstore.BlobStore)
}
