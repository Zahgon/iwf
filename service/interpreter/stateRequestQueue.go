package interpreter

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
)

type StateRequestQueue struct {
	queue []StateRequest
}

func NewStateRequestQueue() *StateRequestQueue { _ = "STUB: not implemented"; return nil }

func NewStateRequestQueueWithResumeRequests(startReqs []iwfidl.StateMovement, resumeReqs map[string]service.StateExecutionResumeInfo) *StateRequestQueue {
	_ = "STUB: not implemented"
	return nil
}

func (srq *StateRequestQueue) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (srq *StateRequestQueue) TakeAll() []StateRequest {
	_ = "STUB: not implemented"
	// copy the whole slice(pointer)
	return nil
}

//reset to empty slice since each iteration will process all current states in the queue

func (srq *StateRequestQueue) GetAllStateStartRequests() []iwfidl.StateMovement {
	_ = "STUB: not implemented"
	return nil
}

func (srq *StateRequestQueue) GetAllStateResumeRequests() []service.StateExecutionResumeInfo {
	_ = "STUB: not implemented"
	return nil
}

func (srq *StateRequestQueue) AddStateStartRequests(reqs []iwfidl.StateMovement) {
	_ = "STUB: not implemented"
	return
}

func (srq *StateRequestQueue) AddSingleStateStartRequest(stateId string, input *iwfidl.EncodedObject, options *iwfidl.WorkflowStateOptions) {
	_ = "STUB: not implemented"
	return
}
