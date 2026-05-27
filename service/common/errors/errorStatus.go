package errors

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
)

type ErrorAndStatus struct {
	StatusCode int
	Error      iwfidl.ErrorResponse
}

func NewErrorAndStatus(statusCode int, subStatus iwfidl.ErrorSubStatus, details string) *ErrorAndStatus {
	_ = "STUB: not implemented"
	return nil
}

func NewErrorAndStatusWithWorkerError(
	statusCode int, subStatus iwfidl.ErrorSubStatus, details string,
	originalWorkerDetails string, originalWorkerErrType string, originalWorkerStatus int32,
) *ErrorAndStatus {
	_ = "STUB: not implemented"
	return nil
}
