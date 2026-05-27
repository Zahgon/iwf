package cont

import (
	"github.com/indeedeng/iwf/service/interpreter/config"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

type ContinueAsNewCounter struct {
	executedStateApis  int32
	signalsReceived    int32
	syncUpdateReceived int32
	triggeredByAPI     bool

	configer *config.WorkflowConfiger
	rootCtx  interfaces.UnifiedContext
	provider interfaces.WorkflowProvider
}

func NewContinueAsCounter(
	configer *config.WorkflowConfiger, rootCtx interfaces.UnifiedContext, provider interfaces.WorkflowProvider,
) *ContinueAsNewCounter {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContinueAsNewCounter) IncExecutedStateExecution(skipStart bool) {
	_ = "STUB: not implemented"
	return
}

func (c *ContinueAsNewCounter) IncSignalsReceived() { _ = "STUB: not implemented"; return }

func (c *ContinueAsNewCounter) IncSyncUpdateReceived() { _ = "STUB: not implemented"; return }

func (c *ContinueAsNewCounter) IsThresholdMet() bool { _ = "STUB: not implemented"; return false }

// Note: when threshold == 0, it means unlimited

func (c *ContinueAsNewCounter) TriggerByAPI() { _ = "STUB: not implemented"; return }
