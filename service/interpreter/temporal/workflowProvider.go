package temporal

import (
	"time"

	"github.com/indeedeng/iwf/service/interpreter/interfaces"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	"go.temporal.io/sdk/workflow"
)

type workflowProvider struct {
	threadCount        int
	pendingThreadNames map[string]int
}

func newTemporalWorkflowProvider() interfaces.WorkflowProvider {
	_ = "STUB: not implemented"
	return *new(interfaces.WorkflowProvider)
}

func (w *workflowProvider) GetBackendType() service.BackendType {
	_ = "STUB: not implemented"
	return *new(service.BackendType)
}

func (w *workflowProvider) NewApplicationError(errType string, details interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) IsApplicationError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *workflowProvider) NewInterpreterContinueAsNewError(
	ctx interfaces.UnifiedContext, input service.InterpreterWorkflowInput,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) UpsertSearchAttributes(
	ctx interfaces.UnifiedContext, attributes map[string]interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) UpsertMemo(ctx interfaces.UnifiedContext, rawMemo map[string]iwfidl.EncodedObject) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) NewTimer(ctx interfaces.UnifiedContext, d time.Duration) interfaces.Future {
	_ = "STUB: not implemented"
	return *new(interfaces.Future)
}

func (w *workflowProvider) GetWorkflowInfo(ctx interfaces.UnifiedContext) interfaces.WorkflowInfo {
	_ = "STUB: not implemented"
	return *new(interfaces.WorkflowInfo)
}

func (w *workflowProvider) GetSearchAttributes(
	ctx interfaces.UnifiedContext, requestedSearchAttributes []iwfidl.SearchAttributeKeyAndType,
) (map[string]iwfidl.SearchAttribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *workflowProvider) SetQueryHandler(
	ctx interfaces.UnifiedContext, queryType string, handler interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) SetRpcUpdateHandler(
	ctx interfaces.UnifiedContext, updateType string, validator interfaces.UnifiedRpcValidator,
	handler interfaces.UnifiedRpcHandler,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) ExtendContextWithValue(
	parent interfaces.UnifiedContext, key string, val interface{},
) interfaces.UnifiedContext {
	_ = "STUB: not implemented"
	return *new(interfaces.UnifiedContext)
}

func (w *workflowProvider) GoNamed(
	ctx interfaces.UnifiedContext, name string, f func(ctx interfaces.UnifiedContext),
) {
	_ = "STUB: not implemented"
	return
}

func (w *workflowProvider) GetPendingThreadNames() map[string]int {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) GetThreadCount() int { _ = "STUB: not implemented"; return 0 }

func (w *workflowProvider) Await(ctx interfaces.UnifiedContext, condition func() bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) WithActivityOptions(
	ctx interfaces.UnifiedContext, options interfaces.ActivityOptions,
) interfaces.UnifiedContext {
	_ = "STUB: not implemented"
	return *new(interfaces.UnifiedContext)
}

// in Temporal, scheduled to close timeout is the timeout include all retries

// support local activity optimization

// set the LA timeout to 7s to make sure the workflow will not need a heartbeat

type futureImpl struct {
	future workflow.Future
}

func (t *futureImpl) IsReady() bool { _ = "STUB: not implemented"; return false }

func (t *futureImpl) Get(ctx interfaces.UnifiedContext, valuePtr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) ExecuteActivity(
	valuePtr interface{}, optimizeByLocalActivity bool,
	ctx interfaces.UnifiedContext, activity interface{}, args ...interface{},
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) ExecuteLocalActivity(
	valuePtr interface{}, ctx interfaces.UnifiedContext, activity interface{}, args ...interface{},
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) Now(ctx interfaces.UnifiedContext) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (w *workflowProvider) Sleep(ctx interfaces.UnifiedContext, d time.Duration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) IsReplaying(ctx interfaces.UnifiedContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *workflowProvider) GetVersion(
	ctx interfaces.UnifiedContext, changeID string, minSupported, maxSupported int,
) int {
	_ = "STUB: not implemented"
	return 0
}

type temporalReceiveChannel struct {
	channel workflow.ReceiveChannel
}

func (t *temporalReceiveChannel) ReceiveAsync(valuePtr interface{}) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (t *temporalReceiveChannel) ReceiveBlocking(ctx interfaces.UnifiedContext, valuePtr interface{}) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (w *workflowProvider) GetSignalChannel(
	ctx interfaces.UnifiedContext, signalName string,
) interfaces.ReceiveChannel {
	_ = "STUB: not implemented"
	return *new(interfaces.ReceiveChannel)
}

func (w *workflowProvider) GetContextValue(ctx interfaces.UnifiedContext, key string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (w *workflowProvider) GetLogger(ctx interfaces.UnifiedContext) interfaces.UnifiedLogger {
	_ = "STUB: not implemented"
	return *new(interfaces.UnifiedLogger)
}

func (w *workflowProvider) GetUnhandledSignalNames(ctx interfaces.UnifiedContext) []string {
	_ = "STUB: not implemented"
	return nil
}
