package interpreter

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/interpreter/interfaces"
)

type PersistenceManager struct {
	dataAttributes   map[string]iwfidl.KeyValue
	searchAttributes map[string]iwfidl.SearchAttribute
	provider         interfaces.WorkflowProvider

	lockedDataAttributesKeys  map[string]bool
	lockedSearchAttributeKeys map[string]bool

	useMemo bool
}

func NewPersistenceManager(
	provider interfaces.WorkflowProvider, initDataAttributes []iwfidl.KeyValue, initSearchAttributes []iwfidl.SearchAttribute,
	useMemo bool,
) *PersistenceManager {
	_ = "STUB: not implemented"
	return nil
}

func RebuildPersistenceManager(
	provider interfaces.WorkflowProvider,
	dolist []iwfidl.KeyValue, salist []iwfidl.SearchAttribute,
	useMemo bool,
) *PersistenceManager {
	_ = "STUB: not implemented"
	return nil
}

// locks will not be carried over during continueAsNew

func (am *PersistenceManager) GetDataAttributesByKey(request service.GetDataAttributesQueryRequest) service.GetDataAttributesQueryResponse {
	_ = "STUB: not implemented"
	return *new(service.GetDataAttributesQueryResponse)
}

func (am *PersistenceManager) LoadSearchAttributes(
	ctx interfaces.UnifiedContext, loadingPolicy *iwfidl.PersistenceLoadingPolicy,
) []iwfidl.SearchAttribute {
	_ = "STUB: not implemented"
	return nil
}

func (am *PersistenceManager) LoadDataAttributes(
	ctx interfaces.UnifiedContext, loadingPolicy *iwfidl.PersistenceLoadingPolicy,
) []iwfidl.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func (am *PersistenceManager) GetAllSearchAttributes() []iwfidl.SearchAttribute {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: using DeterministicKeys so that the JSON snapshot for continueAsNew is stable for pagination
// TODO: we should use DeterministicKeys for every map iteration in interpreter for safety
// https://github.com/indeedeng/iwf/issues/510

func (am *PersistenceManager) GetAllDataAttributes() []iwfidl.KeyValue {
	_ = "STUB: not implemented"
	return nil

	// NOTE: using DeterministicKeys so that the JSON snapshot for continueAsNew is stable for pagination
	// TODO: we should use DeterministicKeys for every map iteration in interpreter for safety
	// https://github.com/indeedeng/iwf/issues/510
}

func (am *PersistenceManager) ProcessUpsertSearchAttribute(
	ctx interfaces.UnifiedContext, attributes []iwfidl.SearchAttribute,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (am *PersistenceManager) ProcessUpsertDataAttribute(ctx interfaces.UnifiedContext, attributes []iwfidl.KeyValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (am *PersistenceManager) CheckDataAndSearchAttributesKeysAreUnlocked(dataAttrKeysToCheck, searchAttrKeysToCheck []string) bool {
	_ = "STUB: not implemented"
	return false
}

func (am *PersistenceManager) checkKeysAreUnlocked(lockedKeys map[string]bool, keysToCheck []string) bool {
	_ = "STUB: not implemented"
	return false
}

func (am *PersistenceManager) awaitAndLockForKeys(ctx interfaces.UnifiedContext, lockedKeys map[string]bool, keysToLock []string) {
	_ = "STUB: not implemented"
	// wait until all keys are not locked
	return
}

// then lock the keys

func (am *PersistenceManager) unlockKeys(lockedKeys map[string]bool, keysToUnlock []string) {
	_ = "STUB: not implemented"
	return
}

func (am *PersistenceManager) UnlockPersistence(
	saPolicy *iwfidl.PersistenceLoadingPolicy, daPolicy *iwfidl.PersistenceLoadingPolicy,
) {
	_ = "STUB: not implemented"
	return
}
