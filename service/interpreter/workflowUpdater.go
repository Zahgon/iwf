package interpreter

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/interpreter/config"
	"github.com/indeedeng/iwf/service/interpreter/cont"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

type WorkflowUpdater struct {
	persistenceManager   *PersistenceManager
	provider             interfaces.WorkflowProvider
	continueAsNewer      *ContinueAsNewer
	continueAsNewCounter *cont.ContinueAsNewCounter
	internalChannel      *InternalChannel
	signalReceiver       *SignalReceiver
	stateRequestQueue    *StateRequestQueue
	configer             *config.WorkflowConfiger
	logger               interfaces.UnifiedLogger
	basicInfo            service.BasicInfo
	globalVersioner      *GlobalVersioner
}

func NewWorkflowUpdater(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, persistenceManager *PersistenceManager,
	stateRequestQueue *StateRequestQueue,
	continueAsNewer *ContinueAsNewer, continueAsNewCounter *cont.ContinueAsNewCounter, configer *config.WorkflowConfiger,
	internalChannel *InternalChannel, signalReceiver *SignalReceiver, basicInfo service.BasicInfo,
	globalVersioner *GlobalVersioner,
) (*WorkflowUpdater, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *WorkflowUpdater) handler(
	ctx interfaces.UnifiedContext, input iwfidl.WorkflowRpcRequest,
) (output *interfaces.HandlerOutput, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *WorkflowUpdater) validator(_ interfaces.UnifiedContext, input iwfidl.WorkflowRpcRequest) error {
	_ = "STUB: not implemented"
	return nil
}
