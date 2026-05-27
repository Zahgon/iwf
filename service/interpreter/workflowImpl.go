package interpreter

import (
	uclient "github.com/indeedeng/iwf/service/client"
	"github.com/indeedeng/iwf/service/interpreter/config"
	"github.com/indeedeng/iwf/service/interpreter/cont"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
)

func InterpreterImpl(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, input service.InterpreterWorkflowInput,
) (output *service.InterpreterWorkflowOutput, retErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// send metrics for the workflow result

// we have stopped upsert here in new versions, because it's done in start workflow request

// The below initialization order should be the same as for non-continueAsNew

// We intentionally set the query handler after the continueAsNew/dumpInternal activity.
// This is to ensure the correctness. If we set the query handler before that,
// the query handler could return empty data (since the loading hasn't completed), which will be incorrect response.
// We would rather return server errors and let the client retry later.

// Note that today different errors could overwrite each other, we only support last one wins. we may use multiError to improve.

// this is for an optimization for StateId Search attribute, see refreshIwfExecutingStateIdSearchAttribute in stateExecutionCounter
// Because it will check totalCurrentlyExecutingCount == 0, so it will also work for continueAsNew case

// it's possible that a workflow is started without any starting state
// it will wait for a new state coming in (by RPC results)

// below was a bug in the older version that workflow didn't continue as new
// but have to keep workflow deterministic

// execute in another thread for parallelism
// state must be passed via parameter https://stackoverflow.com/questions/67263092

// this is the case where stateExecStatus == FailureStateExecutionStatus

// state execution fail should fail the workflow, no more processing

// NOTE: decision is only available on this CompletedStateExecutionStatus

// no return so that it can fall through to call MarkStateExecutionCompleted

// no return so that it can fall through to call MarkStateExecutionCompleted

// finally, mark state completed and may also update system search attribute(IwfExecutingStateIds)

// finally, mark state completed and may also update system search attribute(IwfExecutingStateIds)

// noop for WaitingCommandsStateExecutionStatus, because it means continueAsNew
// end of executing one state
// end loop of executing all states from the queue for one iteration

// The conditions here are quite tricky:
// For !stateRequestQueue.IsEmpty(): We need some condition to wait here because all the state execution are running in different thread.
//    Right after the queue are popped it becomes empty. When it's not empty, it means there are new states to execute pushed into the queue,
//    and it's time to wake up the outer loop to go to next iteration. Alternatively, waiting for all current started in this iteration to complete will also work,
//    but not as efficient as this one because it will take much longer time.
// For errToFailWf != nil || forceCompleteWf: this means we need to close workflow immediately
// For stateExecutionCounter.GetTotalCurrentlyExecutingCount() == 0: this means all the state executions have reach "Dead Ends" so the workflow can complete gracefully without output
// For continueAsNewCounter.IsThresholdMet(): this means workflow need to continueAsNew

// NOTE: drain thread before checking errToFailWf/forceCompleteWf so that we can close the workflow if possible

// this could happen for cancellation

// the outer logic will do the actual continue as new

// end loop until no more state can be executed (dead end)

// we have to drain this again because this can be from non-state cases

// NOTE: This must be the last thing before continueAsNew!!!
// Otherwise, there could be signals unhandled

// after draining signals, there could be some changes
// last fail workflow signal, return the workflow so that we don't carry over the fail request

// if it is empty and no stateExecutionsToResume and request a graceful complete just complete the loop
// so that we don't carry over shouldGracefulComplete

// last update config, do it here because we use input to carry over config, not continueAsNewer query
// update config to the latest before continueAsNew to carry over

// nix the unused data

// end main loop

// gracefully complete workflow when all states are executed to dead ends

func checkClosingWorkflow(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, versioner *GlobalVersioner, decision *iwfidl.StateDecision,
	currentStateId, currentStateExeId string,
	internalChannel *InternalChannel, signalReceiver *SignalReceiver,
) (canGoNext, gracefulComplete, forceComplete, forceFail bool, completeOutput *iwfidl.StateCompletionOutput, err error) {
	_ = "STUB: not implemented"
	return false, false, false, false, nil, nil
}

// trigger a signal draining so that all the signal/internal channel messages are processed

// Messages of internal channels could be published via State executions, within the same workflow task.
// If we don't do any draining and process them, the conditional completion could lose the messages

// condition is met, force complete the workflow

// legacy to keep compatibility for old code that use empty decision as graceful complete

// Illegal decision

func DrainReceivedButUnprocessedInternalChannelsFromStateApis(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, versioner *GlobalVersioner,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Just yield, by waiting on an empty lambda, nothing else.
// It will let other workflow threads/coroutines to run.
// This will drain the messages published from state APIs.
// NOTE that this is extremely tricky in Cadence/Temporal programming model.
// Read more: https://stackoverflow.com/questions/71356668/how-does-multi-threading-works-in-cadence-temporal-workflow
//https://docs.temporal.io/encyclopedia/go-sdk-multithreading

func processStateExecution(
	ctx interfaces.UnifiedContext,
	provider interfaces.WorkflowProvider,
	globalVersioner *GlobalVersioner,
	basicInfo service.BasicInfo,
	stateReq StateRequest,
	stateExeId string,
	persistenceManager *PersistenceManager,
	interStateChannel *InternalChannel,
	signalReceiver *SignalReceiver,
	timerProcessor interfaces.TimerProcessor,
	continueAsNewer *ContinueAsNewer,
	continueAsNewCounter *cont.ContinueAsNewCounter,
	configer *config.WorkflowConfiger,
	shouldSendSignalOnCompletion bool,
) (*iwfidl.StateDecision, service.StateExecutionStatus, error) {
	_ = "STUB: not implemented"
	return nil, *new(service.StateExecutionStatus), nil
}

// skip the completed timers(from continueAsNew)

//Start timer in a new thread

// Note that commandReqDoneOrCanceled is needed for two cases:
// 1. will be true when trigger type of the commandReq is completed(e.g. AnyCommandCompleted) so we don't need to wait for all commands. Returning the thread to avoid thread leakage.
// 2. will be true to cancel the wait for unblocking continueAsNew(continueAsNew will wait for all threads to complete)

// skip completed signal(from continueAsNew)

//Process signal in new thread

// Note that commandReqDoneOrCanceled is needed for two cases:
// 1. will be true when trigger type of the commandReq is completed(e.g. AnyCommandCompleted) so we don't need to wait for all commands. Returning the thread to avoid thread leakage.
// 2. will be true to cancel the wait for unblocking continueAsNew(continueAsNew will wait for all threads to complete)

// skip completed interStateChannelCommand(from continueAsNew)

//Process interstate channel command in a new thread.

// Note that commandReqDoneOrCanceled is needed for two cases:
// 1. will be true when trigger type of the commandReq is completed(e.g. AnyCommandCompleted) so we don't need to wait for all commands. Returning the thread to avoid thread leakage.
// 2. will be true to cancel the wait for unblocking continueAsNew(continueAsNew will wait for all threads to complete)

//Passing a map of references of completed or soon to be completed commands (once the above threads are complete) and the state execution variables to the continueAsNewer.
//After this method completes and if continueAsNewCounter.IsThresholdMet() is true, this snapshot will be used to start a new continueAsNew workflow while preserving the state of the workflow at the end of this method.
//This snapshot is also used to query the workflow state, which can be done at anytime.

// Wait for decider trigger (ANY/ALL command completed) OR continue-as-new threshold

//This variable tells all command threads to stop waiting and exit, even if their specific command has not been completed.
//In both cases, the trigger condition has been met or the continue-as-new threshold has been reached we want the above command threads to stop waiting.

// Wait for command threads to drain. After the command request await completes, command threads
// may still be in the process of storing retrieved data into completedXXXCmds maps.
// We must wait for these threads to finish before assembling command results, otherwise
// retrieved data will be lost (the thread retrieved it but never stored it before we return the maps).
// We only wait for threads of commands that currently have data or has been canceled or fired.
// A thread that doesn't have data can be canceled when commandReqDoneOrCanceled is set to true. This preserves ANY_COMMAND_COMPLETED semantics.

// this means continueAsNewCounter.IsThresholdMet == true
// not using continueAsNewCounter.IsThresholdMet because deciderTrigger is higher prioritized
// it won't continueAsNew in those cases 1. start Api fail with proceed policy, 2. empty commands, 3. both commands and continueAsNew are met

func invokeStateExecute(
	ctx interfaces.UnifiedContext,
	provider interfaces.WorkflowProvider,
	basicInfo service.BasicInfo,
	state iwfidl.StateMovement,
	stateExeId string,
	persistenceManager *PersistenceManager,
	interStateChannel *InternalChannel,
	executionContext iwfidl.Context,
	commandRes *iwfidl.CommandResults,
	continueAsNewer *ContinueAsNewer,
	configer *config.WorkflowConfiger,
	executeApi interface{},
	stateExecutionLocal []iwfidl.KeyValue,
	shouldSendSignalOnCompletion bool,
) (*iwfidl.StateDecision, service.StateExecutionStatus, error) {
	_ = "STUB: not implemented"
	return nil, *new(service.StateExecutionStatus), nil
}

// NOTE: here uses NOT IsReplaying to signalWithStart, to save an activity for this operation
// this is not a problem because the signalWithStart will be very fast and highly available

// signalWithStart with legacy workflowId (containing parent workflowId)

// signalWithStart with new workflowId (containing firstRunId)

// Start WaitForStateCompletionWorkflow with a new name to ensure smooth transition

func signalWithStart(unifiedClient uclient.UnifiedClient, workflowId string) {
	_ = "STUB: not implemented"
	return
}

// timeout doesn't matter here as it will complete immediate with the signal

// WorkflowAlreadyStartedError is returned when the started workflow is closed and the signal is not sent
// panic will let the workflow task will retry until the signal is sent

func shouldProceedOnStartApiError(state iwfidl.StateMovement) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldProceedOnExecuteApiError(state iwfidl.StateMovement) bool {
	_ = "STUB: not implemented"
	return false
}

func convertStateApiActivityError(provider interfaces.WorkflowProvider, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func getCommandThreadName(prefix string, stateExecId, cmdId string, idx int) string {
	_ = "STUB: not implemented"
	return ""
}

func createUserWorkflowError(provider interfaces.WorkflowProvider, message string) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitForStateCompletionWorkflowImpl(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider,
) (*service.WaitForStateCompletionWorkflowOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BlobStoreCleanup(
	ctx interfaces.UnifiedContext, provider interfaces.WorkflowProvider, storeId string,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
