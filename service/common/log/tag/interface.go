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

	"go.uber.org/zap"
)

// Tag is the interface for logging system
type Tag struct {
	// keep this field private
	field zap.Field
}

// Field returns a zap field
func (t *Tag) Field() zap.Field { _ = "STUB: not implemented"; return *new(zap.Field) }

func newStringTag(key string, value string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

func newInt64(key string, value int64) Tag { _ = "STUB: not implemented"; return *new(Tag) }

func newInt(key string, value int) Tag { _ = "STUB: not implemented"; return *new(Tag) }

func newInt32(key string, value int32) Tag { _ = "STUB: not implemented"; return *new(Tag) }

func newBoolTag(key string, value bool) Tag { _ = "STUB: not implemented"; return *new(Tag) }

func newErrorTag(key string, value error) Tag {
	_ = "STUB: not implemented"
	// NOTE zap already chosen "error" as key
	return *new(Tag)
}

func newDurationTag(key string, value time.Duration) Tag {
	_ = "STUB: not implemented"
	return *new(Tag)
}

func newTimeTag(key string, value time.Time) Tag { _ = "STUB: not implemented"; return *new(Tag) }

func newObjectTag(key string, value interface{}) Tag { _ = "STUB: not implemented"; return *new(Tag) }

func newPredefinedStringTag(key string, value string) Tag {
	_ = "STUB: not implemented"
	return *new(Tag)
}

func newPredefinedDynamicTag(key string, value interface{}) Tag {
	_ = "STUB: not implemented"
	return *new(Tag)
}
