package utils

import (
	"context"
	"net/http"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	defaultMaxApiTimeoutSeconds = 60
)

func MergeStringSlice(first, second []string) []string { _ = "STUB: not implemented"; return nil }

func MergeMap(first map[string]interface{}, second map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func TrimRpcTimeoutSeconds(ctx context.Context, req iwfidl.WorkflowRpcRequest) int32 {
	_ = "STUB: not implemented"
	return 0
}

func TrimContextByTimeoutWithCappedDDL(parent context.Context, reqWaitSeconds *int32, configuredMaxSeconds int64) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// then capped by context

func CheckHttpError(err error, httpResp *http.Response) bool {
	_ = "STUB: not implemented"
	return false
}

func ToNanoSeconds(e *timestamppb.Timestamp) int64 { _ = "STUB: not implemented"; return 0 }

func GetWorkflowIdForWaitForStateExecution(parentId string, stateExeId *string, waitForKey *string, stateId *string) string {
	_ = "STUB: not implemented"
	return ""
}
