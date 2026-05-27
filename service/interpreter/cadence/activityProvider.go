package cadence

import (
	"context"

	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

type activityProvider struct{}

func init() {
	interfaces.RegisterActivityProvider(service.BackendTypeCadence, &activityProvider{})
}

func (a *activityProvider) NewApplicationError(errType string, details interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *activityProvider) GetLogger(ctx context.Context) interfaces.UnifiedLogger {
	_ = "STUB: not implemented"
	return *new(interfaces.UnifiedLogger)
}

func (a *activityProvider) GetActivityInfo(ctx context.Context) interfaces.ActivityInfo {
	_ = "STUB: not implemented"
	return *new(interfaces.ActivityInfo)
}

// NOTE increase by one to match Temporal
// TODO cadence doesn't support this yet

func (a *activityProvider) RecordHeartbeat(ctx context.Context, details ...interface{}) {
	_ = "STUB: not implemented"
	return
}
