package s3_upsert_data_objects

import (
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

/**
 * This test workflow has 2 states, testing S3 upsert data objects functionality.
 *
 * State1:
 *		- WaitUntil method does nothing
 *      - Execute method upserts large data objects that should go to S3, then transitions to State2
 *
 * State2:
 *		- WaitUntil method validates it receives the upserted data objects from S3
 *      - Execute method completes workflow
 */
const (
	WorkflowType      = "s3-upsert-data-objects"
	State1            = "S1"
	State2            = "S2"
	TestDataObjKey1   = "large_obj1"
	TestDataObjKey2   = "large_obj2"
	TestDataObjKey3   = "small_obj3"
	LargeDataContent1 = "this_is_a_large_data_content_that_should_be_stored_in_s3_for_upsert_testing_purposes_with_more_than_10_characters"
	LargeDataContent2 = "another_large_data_content_for_second_upserted_object_that_exceeds_the_s3_threshold_for_external_storage_testing"
	SmallDataContent3 = "small"
)

type handler struct {
	invokeHistory sync.Map
	invokeData    sync.Map
}

func NewHandler() *handler { _ = "STUB: not implemented"; return nil }

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Increment invoke count

// State1 waitUntil doesn't need to do anything

// State2 waitUntil should validate data objects received from State1's upsert

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Increment invoke count

// State1 Execute: Upsert large data objects that should go to S3

// Large - should go to S3

// Large - should go to S3

// Small - should stay in memory

// Transition to State2

// State2 Execute: Complete workflow

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Merge both maps
