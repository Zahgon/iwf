package interpreter

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
)

type StateRequest struct {
	stateStartRequest  iwfidl.StateMovement
	isResumeRequest    bool
	stateResumeRequest service.StateExecutionResumeInfo
}

func NewStateStartRequest(movement iwfidl.StateMovement) StateRequest {
	_ = "STUB: not implemented"
	return *new(StateRequest)
}

func NewStateResumeRequest(resumeRequest service.StateExecutionResumeInfo) StateRequest {
	_ = "STUB: not implemented"
	return *new(StateRequest)
}

func (sq StateRequest) GetStateStartRequest() iwfidl.StateMovement {
	_ = "STUB: not implemented"
	return *new(iwfidl.StateMovement)
}

func (sq StateRequest) GetStateResumeRequest() service.StateExecutionResumeInfo {
	_ = "STUB: not implemented"
	return *new(service.StateExecutionResumeInfo)
}

func (sq StateRequest) IsResumeRequest() bool { _ = "STUB: not implemented"; return false }

func (sq StateRequest) GetStateMovement() iwfidl.StateMovement {
	_ = "STUB: not implemented"
	return *new(iwfidl.StateMovement)
}

func (sq StateRequest) GetStateId() string { _ = "STUB: not implemented"; return "" }
