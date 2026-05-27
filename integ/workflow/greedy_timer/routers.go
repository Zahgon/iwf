package greedy_timer

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/integ/workflow/common"
)

/*
*
This workflow will accept an array of integers representing durations and execute a state that waits on a timer corresponding to each duration provided
*/
const (
	WorkflowType       = "greedy_timer"
	ScheduleTimerState = "schedule"
	SubmitDurationsRPC = "submitDurationsRPC"
)

type handler struct {
	invokeHistory sync.Map
	invokeData    sync.Map
}

type Input struct {
	Durations []int64 `json:"durations"`
}

func NewHandler() common.WorkflowHandlerWithRpc {
	_ = "STUB: not implemented"
	return *new(common.WorkflowHandlerWithRpc)
}

func (h *handler) ApiV1WorkflowWorkerRpc(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}
