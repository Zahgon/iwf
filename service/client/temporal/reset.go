package temporal

import (
	"context"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/converter"
)

func getResetEventIDByType(ctx context.Context, resetType iwfidl.WorkflowResetType,
	namespace, wid, rid string,
	frontendClient workflowservice.WorkflowServiceClient, converter converter.DataConverter,
	historyEventId int32, earliestHistoryTimeStr string, stateId, stateExecutionId string,
) (resetBaseRunID string, workflowTaskFinishID int64, err error) {
	_ = "STUB: not implemented"
	// default to the same runID
	return "", 0, nil
}

func getFirstWorkflowTaskEventID(ctx context.Context, namespace, wid, rid string, frontendClient workflowservice.WorkflowServiceClient) (resetBaseRunID string, workflowTaskEventID int64, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func getEarliestDecisionEventID(
	ctx context.Context,
	namespace string, wid string,
	rid string, earliestTime int64,
	frontendClient workflowservice.WorkflowServiceClient,
) (decisionFinishID int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getDecisionEventIDByStateOrStateExecutionId(
	ctx context.Context,
	namespace string, wid string,
	rid string, stateId, stateExecutionId string,
	frontendClient workflowservice.WorkflowServiceClient, converter converter.DataConverter,
) (decisionFinishID int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//TODO: Add check for local activity. (IWF-403)

func composeErrorWithMessage(msg string, err error) error { _ = "STUB: not implemented"; return nil }
