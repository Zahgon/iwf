package interpreter

import (
	"github.com/indeedeng/iwf/service/interpreter/config"
	"github.com/indeedeng/iwf/service/interpreter/cont"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"

	"github.com/indeedeng/iwf/gen/iwfidl"
)

type SignalReceiver struct {
	// key is channel name
	receivedSignals            map[string][]*iwfidl.EncodedObject
	failWorkflowByClient       bool
	reasonFailWorkflowByClient *string
	provider                   interfaces.WorkflowProvider
	timerProcessor             interfaces.TimerProcessor
	workflowConfiger           *config.WorkflowConfiger
	interStateChannel          *InternalChannel
	stateRequestQueue          *StateRequestQueue
	persistenceManager         *PersistenceManager
}

func NewSignalReceiver(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, interStateChannel *InternalChannel,
	stateRequestQueue *StateRequestQueue,
	persistenceManager *PersistenceManager, tp interfaces.TimerProcessor, continueAsNewCounter *cont.ContinueAsNewCounter,
	workflowConfiger *config.WorkflowConfiger,
	initReceivedSignals map[string][]*iwfidl.EncodedObject,
) *SignalReceiver {
	_ = "STUB: not implemented"
	return nil
}

//The thread waits until a FailWorkflowSignalChannelName signal has been
//received or a continueAsNew run is triggered. When a signal has been received it sets
//SignalReceiver.failWorkflowByClient to true and sets SignalReceiver.reasonFailWorkflowByClient to the reason
//given in the signal's value. If continueIsNew is triggered, the thread completes after all signals have been processed.

// NOTE: continueAsNew will wait for all threads to complete, so we must stop this thread for continueAsNew when no more signals to process

// NOTE: continueAsNew will wait for all threads to complete, so we must stop this thread for continueAsNew when no more signals to process

//The thread waits until a SkipTimerSignalChannelName signal has been
//received or a continueAsNew run is triggered. When a signal has been received it skips the specific timer
//described in the signal's value. If continueIsNew is triggered, the thread completes after all signals have been processed.

// break the loop to prevent goroutine leakage

// NOTE: continueAsNew will wait for all threads to complete, so we must stop this thread for continueAsNew when no more signals to process

//The thread waits until a UpdateConfigSignalChannelName signal has been
//received or a continueAsNew run is triggered. When a signal has been received it updates the workflow config
//defined in the signal's value. If continueIsNew is triggered, the thread completes after all signals have been processed.

// break the loop to prevent goroutine leakage

// NOTE: continueAsNew will wait for all threads to complete, so we must stop this thread for continueAsNew when no more signals to process

//The thread waits until a TriggerContinueAsNewSignalChannelName signal has
//been received or a continueAsNew run is triggered. When a signal has been received it triggers a continueAsNew run.
//Since this thread is triggering a continueAsNew run it doesn't need to wait for signals to drain from the channel.

// NOTE: unlike other signal channels, this one doesn't need to drain during continueAsNew
// because if there is a continueAsNew, this signal is not needed anymore

//The thread waits until a ExecuteRpcSignalChannelName signal has been
//received or a continueAsNew run is triggered. When a signal has been received it upserts data objects
//(if they exist in the signal value), upserts search attributes (if they exist in the signal value),
//and/or publishes a message to an internal channel (if InterStateChannelPublishing is set in the signal value).
//If continueIsNew is triggered, the thread completes after all signals have been processed.

// break the loop to prevent goroutine leakage

// NOTE: continueAsNew will wait for all threads to complete, so we must stop this thread for continueAsNew when no more signals to process

//The thread waits until a signal has been received that is not an IWF
//system signal name or a continueAsNew run is triggered. When a signal has been received it processes the
//external signal. If continueIsNew is triggered, the thread completes after all signals have been processed.

// skip this because it will be processed in a different thread

// NOTE: continueAsNew will wait for all threads to complete, so we must stop this thread for continueAsNew when no more signals to process

func (sr *SignalReceiver) receiveSignal(ctx interfaces.UnifiedContext, sigName string) {
	_ = "STUB: not implemented"
	return
}

func (sr *SignalReceiver) HasSignal(channelName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (sr *SignalReceiver) Retrieve(channelName string) *iwfidl.EncodedObject {
	_ = "STUB: not implemented"
	return nil
}

func (sr *SignalReceiver) GetAllReceived() map[string][]*iwfidl.EncodedObject {
	_ = "STUB: not implemented"
	return nil
}

func (sr *SignalReceiver) GetInfos() map[string]iwfidl.ChannelInfo {
	_ = "STUB: not implemented"
	return nil
}

// DrainAllReceivedButUnprocessedSignals will process all the signals that are received but not processed in the current
// workflow task.
// There are two cases this is needed:
// 1. ContinueAsNew:
// retrieve signals that after signal handler threads are stopped,
// so that the signals can be carried over to next run by continueAsNew.
// This includes both regular user signals and system signals
// 2. Conditional close/complete workflow on signal/internal channel:
// retrieve all signal/internal channel messages before checking the signal/internal channels
func (sr *SignalReceiver) DrainAllReceivedButUnprocessedSignals(ctx interfaces.UnifiedContext) {
	_ = "STUB: not implemented"
	return
}

// ignore invalid system signals because we can't process it

func (sr *SignalReceiver) IsFailWorkflowRequested() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
