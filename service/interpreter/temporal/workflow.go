package temporal

import (
	"github.com/indeedeng/iwf/service"
	"go.temporal.io/sdk/workflow"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

func Interpreter(ctx workflow.Context, input service.InterpreterWorkflowInput) (*service.InterpreterWorkflowOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WaitforStateCompletionWorkflow(ctx workflow.Context) (*service.WaitForStateCompletionWorkflowOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BlobStoreCleanup(ctx workflow.Context, storeId string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
