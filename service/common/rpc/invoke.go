package rpc

import (
	"context"

	"net/http"

	"github.com/indeedeng/iwf/config"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/common/blobstore"
	"github.com/indeedeng/iwf/service/common/errors"
)

func InvokeWorkerRpc(
	ctx context.Context, rpcPrep *service.PrepareRpcQueryResponse, req iwfidl.WorkflowRpcRequest, apiMaxSeconds int64, blobStore blobstore.BlobStore, externalStorageConfig config.ExternalStorageConfig,
) (*iwfidl.WorkflowWorkerRpcResponse, *errors.ErrorAndStatus) {
	_ = "STUB: not implemented"
	return nil, nil
}

// invoke worker rpc

// creating empty maps for signalChannelInfos & internalChannelInfos instead of passing in nils
// using nil causes problems when converting to map model defined with OpenAPI

// TODO this need more work in workflow to support

func handleWorkerRpcResponseError(err error, httpResp *http.Response) *errors.ErrorAndStatus {
	_ = "STUB: not implemented"
	return nil
}
