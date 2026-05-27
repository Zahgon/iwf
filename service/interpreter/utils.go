package interpreter

import (
	"golang.org/x/exp/constraints"
)

func caller(skip int) string { _ = "STUB: not implemented"; return "" }

func LastCaller() string {
	_ = "STUB: not implemented"

	// DeterministicKeys returns the keys of a map in deterministic (sorted) order. To be used in for
	// loops in workflows for deterministic iteration.
	return ""
}

func DeterministicKeys[K constraints.Ordered, V any](m map[K]V) []K {
	_ = "STUB: not implemented"
	// copy from https://github.com/temporalio/sdk-go/blob/7828e06cf517dd2d27881a33efaaf4ff985f2e14/workflow/workflow.go#L787
	// and example usage https://github.com/temporalio/samples-go/blob/c69dc0bacc78163a50465c6f80aa678739673a4d/safe_message_handler/workflow.go#L119
	return nil
}
