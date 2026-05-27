package mapper

import (
	"github.com/indeedeng/iwf/gen/iwfidl"
	"go.temporal.io/api/common/v1"
	"go.uber.org/cadence/.gen/go/shared"
)

func MapToInternalSearchAttributes(attributes []iwfidl.SearchAttribute) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MapCadenceToIwfSearchAttributes(searchAttributes *shared.SearchAttributes, requestedSearchAttributes []iwfidl.SearchAttributeKeyAndType) (map[string]iwfidl.SearchAttribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// when SAs are nil, skip calling mapToIwfSearchAttribute

func MapTemporalToIwfSearchAttributes(searchAttributes *common.SearchAttributes, requestedSearchAttributes []iwfidl.SearchAttributeKeyAndType) (map[string]iwfidl.SearchAttribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: Temporal require search attributes always use default data converter, so we don't need to use the customized one

// TODO we should also call UseNumber here for JSON decoder for Temporal
// see https://github.com/temporalio/sdk-go/issues/942
// when SAs are nil, skip calling mapToIwfSearchAttribute

func mapToIwfSearchAttribute(key string, valueType iwfidl.SearchAttributeValueType, object interface{}, useNumber bool) (*iwfidl.SearchAttribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for DATETIME it will be like 2022-12-27T20:00:24.338155843Z
