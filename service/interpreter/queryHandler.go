package interpreter

import (
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/interpreter/config"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

func SetQueryHandlers(
	ctx interfaces.UnifiedContext,
	provider interfaces.WorkflowProvider,
	timerProcessor interfaces.TimerProcessor,
	persistenceManager *PersistenceManager,
	internalChannel *InternalChannel,
	signalReceiver *SignalReceiver,
	continueAsNewer *ContinueAsNewer,
	workflowConfiger *config.WorkflowConfiger,
	basicInfo service.BasicInfo,
) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO use firstRunId instead
