// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tag

import (
	"time"
)

// LoggingCallAtKey is reserved tag
const LoggingCallAtKey = "logging-call-at"

// All logging tags are defined in this file.
// To help finding available tags, we recommend that all tags to be categorized and placed in the corresponding section.
// We currently have those categories:
//   0. Common tags that can't be categorized(or belong to more than one)
//   1. Workflow: these tags are information that are useful to our customer, like workflow-id/run-id/task-list/...
//   2. System : these tags are internal information which usually cannot be understood by our customers,

///////////////////  Common tags defined here ///////////////////

// Error returns tag for Error
func Error(err error) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Timestamp returns tag for Timestamp
func Timestamp(timestamp time.Time) Tag { _ = "STUB: not implemented"; return *new(Tag) }

///////////////////  Workflow tags defined here: ( wf is short for workflow) ///////////////////

// WorkflowAction returns tag for WorkflowAction
func workflowAction(action string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// general

// Service returns tag for Service
func Service(sv string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowError returns tag for WorkflowError
func WorkflowError(error error) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowTimeoutType returns tag for WorkflowTimeoutType
func WorkflowTimeoutType(timeoutType int64) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowID returns tag for WorkflowID
func WorkflowID(workflowID string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// StatusCode returns tag for StatusCode
func StatusCode(code int) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// SubStatus returns tag for SubStatus
func SubStatus(status string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowType returns tag for WorkflowType
func WorkflowType(wfType string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowState returns tag for WorkflowState
func WorkflowState(s int) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowRunID returns tag for WorkflowRunID
func WorkflowRunID(runID string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowResetBaseRunID returns tag for WorkflowResetBaseRunID
func WorkflowResetBaseRunID(runID string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowResetNewRunID returns tag for WorkflowResetNewRunID
func WorkflowResetNewRunID(runID string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowBinaryChecksum returns tag for WorkflowBinaryChecksum
func WorkflowBinaryChecksum(cs string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// WorkflowActivityID returns tag for WorkflowActivityID
func WorkflowActivityID(id string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// OperationName returns tag for OperationName
func OperationName(operationName string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// history event ID related

// WorkflowEventID returns tag for WorkflowEventID
func WorkflowEventID(eventID int64) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Address return tag for Address
func Address(ad string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Env return tag for runtime environment
func Env(env string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Key returns tag for Key
func Key(k string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Name returns tag for Name
func Name(k string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Value returns tag for Value
func Value(v interface{}) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// ValueType returns tag for ValueType
func ValueType(v interface{}) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// DefaultValue returns tag for DefaultValue
func DefaultValue(v interface{}) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Port returns tag for Port
func Port(p int) Tag {
	_ = "STUB: not implemented"
	return *

	// Counter returns tag for Counter
	new(Tag)
}

func Counter(c int) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Number returns tag for Number
func Number(n int64) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// NextNumber returns tag for NextNumber
func NextNumber(n int64) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Bool returns tag for Bool
func Bool(b bool) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// SysStackTrace returns tag for SysStackTrace
func SysStackTrace(stackTrace string) Tag { _ = "STUB: not implemented"; return *new(Tag) }
