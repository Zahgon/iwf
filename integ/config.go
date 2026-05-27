package integ

import (
	"github.com/indeedeng/iwf/config"
)

const testWorkflowServerPort = "9714"
const testIwfServerPort = "9715"

func createTestConfig(testCfg IwfServiceTestConfig) config.Config {
	_ = "STUB: not implemented"
	return *new(config.Config)
}

// use 12 so that we can test it in the waiting test
