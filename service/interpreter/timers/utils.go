package timers

import (
	"time"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
)

func removeElement(s []service.StaleSkipTimerSignal, i int) []service.StaleSkipTimerSignal {
	_ = "STUB: not implemented"
	return nil
}

// FixTimerCommandFromActivityOutput converts the durationSeconds to firingUnixTimestampSeconds
// doing it right after the activity output so that we don't need to worry about the time drift after continueAsNew
func FixTimerCommandFromActivityOutput(now time.Time, request iwfidl.CommandRequest) iwfidl.CommandRequest {
	_ = "STUB: not implemented"
	return *new(iwfidl.CommandRequest)
}
