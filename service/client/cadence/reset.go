package cadence

import (
	"context"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/cadence/encoded"
)

func getResetIDsByType(
	ctx context.Context,
	resetType iwfidl.WorkflowResetType,
	domain, wid, rid string,
	frontendClient workflowserviceclient.Interface, converter encoded.DataConverter,
	historyEventId int32, earliestHistoryTimeStr string, stateId, stateExecutionId string,
) (resetBaseRunID string, decisionFinishID int64, err error) {
	_ = "STUB: not implemented"
	// default to the same runID
	return "", 0, nil
}

func getFirstDecisionTaskByType(
	ctx context.Context,
	domain string,
	workflowID string,
	runID string,
	frontendClient workflowserviceclient.Interface,
	decisionType shared.EventType,
) (decisionFinishID int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getEarliestDecisionID(
	ctx context.Context,
	domain string, wid string,
	rid string, earliestTime int64,
	frontendClient workflowserviceclient.Interface,
) (decisionFinishID int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getDecisionEventIDByStateOrStateExecutionId(
	ctx context.Context,
	domain string, wid string,
	rid string, stateId, stateExecutionId string,
	frontendClient workflowserviceclient.Interface,
	converter encoded.DataConverter,
) (decisionFinishID int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//TODO: Add check for local activity. (IWF-403)

func composeErrorWithMessage(msg string, err error) error { _ = "STUB: not implemented"; return nil }
