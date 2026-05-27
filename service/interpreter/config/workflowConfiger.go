package config

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
)

type WorkflowConfiger struct {
	config iwfidl.WorkflowConfig
}

func NewWorkflowConfiger(config iwfidl.WorkflowConfig) *WorkflowConfiger {
	_ = "STUB: not implemented"
	return nil
}

func (wc *WorkflowConfiger) Get() iwfidl.WorkflowConfig {
	_ = "STUB: not implemented"
	return *new(iwfidl.WorkflowConfig)
}

func (wc *WorkflowConfiger) ShouldOptimizeActivity() bool { _ = "STUB: not implemented"; return false }

func (wc *WorkflowConfiger) UpdateByAPI(config iwfidl.WorkflowConfig) {
	_ = "STUB: not implemented"
	return
}
