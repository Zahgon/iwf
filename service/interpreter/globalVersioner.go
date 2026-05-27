package interpreter

import (
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

const globalChangeId = "global"

// StartingVersionUsingGlobalVersioning First global version
const StartingVersionUsingGlobalVersioning = 1

// StartingVersionOptimizedUpsertSearchAttribute Optimized upserting SAs
const StartingVersionOptimizedUpsertSearchAttribute = 2

// StartingVersionRenamedStateApi Renamed state API
// see: https://github.com/indeedeng/iwf/pull/242/files
const StartingVersionRenamedStateApi = 3

// StartingVersionContinueAsNewOnNoStates Fix ContinueAsNew bug
const StartingVersionContinueAsNewOnNoStates = 4

// StartingVersionTemporal26SDK Upgraded Temporal SDK version which brought changes to update handler
// see: https://github.com/indeedeng/iwf/releases/tag/v1.11.0
const StartingVersionTemporal26SDK = 5

// StartingVersionExecutingStateIdMode Changed default rule of upserting SAs
const StartingVersionExecutingStateIdMode = 6

// StartingVersionNoIwfGlobalVersionSearchAttribute Removed upserting IwfGlobalWorkflowVersion SA
const StartingVersionNoIwfGlobalVersionSearchAttribute = 7

// StartingVersionYieldOnConditionalComplete Bug fix to where published messages could be lost
const StartingVersionYieldOnConditionalComplete = 8

// SyncUpdateRPCUseLocalActivity Always use local activities for sync update based RPC
const SyncUpdateRPCUseLocalActivity = 9

// StartingVersionWaitingCommandThreads waits for all command threads to complete before taking a snapshot.
// This ensures that commands don't get lost during continueAsNew operations.
const StartingVersionWaitingCommandThreads = 10

const MaxOfAllVersions = StartingVersionWaitingCommandThreads

// GlobalVersioner see https://stackoverflow.com/questions/73941723/what-is-a-good-way-pattern-to-use-temporal-cadence-versioning-api
type GlobalVersioner struct {
	workflowProvider interfaces.WorkflowProvider
	ctx              interfaces.UnifiedContext
	version          int
}

func NewGlobalVersioner(
	workflowProvider interfaces.WorkflowProvider, ctx interfaces.UnifiedContext,
) (*GlobalVersioner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// methods checking version number

func (p *GlobalVersioner) IsAfterVersionOfContinueAsNewOnNoStates() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) IsAfterVersionOfUsingGlobalVersioning() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) IsAfterVersionOfOptimizedUpsertSearchAttribute() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) IsAfterVersionOfExecutingStateIdMode() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) IsAfterVersionOfRenamedStateApi() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) IsAfterVersionOfTemporal26SDK() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) IsAfterVersionOfNoIwfGlobalVersionSearchAttribute() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) IsAfterVersionOfYieldOnConditionalComplete() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) IsAfterVersionOfSyncUpdateRPCUseLocalActivity() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) IsAfterVersionOfWaitingCommandThreads() bool {
	_ = "STUB: not implemented"
	return false
}

// methods checking feature/functionality availability

func (p *GlobalVersioner) IsUsingGlobalVersionSearchAttribute() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *GlobalVersioner) UpsertGlobalVersionSearchAttribute() error {
	_ = "STUB: not implemented"
	return nil
}

// Note that there was bug in Cadence SDK may cause concurrent writes hence we never upsert for Cadence
// https://github.com/uber-go/cadence-client/issues/1198
