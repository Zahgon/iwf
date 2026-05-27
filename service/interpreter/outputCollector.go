package interpreter

import "github.com/indeedeng/iwf/gen/iwfidl"

type OutputCollector struct {
	outputs []iwfidl.StateCompletionOutput
}

func NewOutputCollector(initOutputs []iwfidl.StateCompletionOutput) *OutputCollector {
	_ = "STUB: not implemented"
	return nil
}

func (o *OutputCollector) Add(output iwfidl.StateCompletionOutput) {
	_ = "STUB: not implemented"
	return
}

func (o *OutputCollector) GetAll() []iwfidl.StateCompletionOutput {
	_ = "STUB: not implemented"
	return nil
}
