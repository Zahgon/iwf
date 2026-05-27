package event

import "github.com/indeedeng/iwf/gen/iwfidl"

// The implementation must be lightweight, reliable and fast (less than 1s)
type HandleEventFunc func(event iwfidl.IwfEvent)

var Handle HandleEventFunc = DefaultHandleEventFunc

func SetHandleEventFunc(handler HandleEventFunc) { _ = "STUB: not implemented"; return }

func DefaultHandleEventFunc(event iwfidl.IwfEvent) {
	_ = "STUB: not implemented"
	// Noop by default
	return
}
