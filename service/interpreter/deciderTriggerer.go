package interpreter

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
)

func IsDeciderTriggerConditionMet(
	commandReq iwfidl.CommandRequest,
	completedTimerCmds map[int]service.InternalTimerStatus,
	completedSignalCmds map[int]*iwfidl.EncodedObject,
	completedInterStateChannelCmds map[int]*iwfidl.EncodedObject,
) bool {
	_ = "STUB: not implemented"
	return false
}
