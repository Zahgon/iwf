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

package loggerimpl

import (
	"github.com/indeedeng/iwf/service/common/log"
	"github.com/indeedeng/iwf/service/common/log/tag"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type loggerImpl struct {
	zapLogger *zap.Logger
	skip      int
}

const (
	skipForDefaultLogger = 3
	// we put a default message when it is empty so that the log can be searchable/filterable
	defaultMsgForEmpty = "none"
)

// NewNopLogger returns a no-op logger
func NewNopLogger() log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

// NewLoggerForTest is a helper to create new development logger in unit test
func NewLoggerForTest(s suite.Suite) log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }

// NewDevelopment returns a logger at debug level and log into STDERR
func NewDevelopment() (log.Logger, error) { _ = "STUB: not implemented"; return *new(log.Logger), nil }

// NewLogger returns a new logger
func NewLogger(zapLogger *zap.Logger) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

func caller(skip int) string { _ = "STUB: not implemented"; return "" }

func (lg *loggerImpl) buildFieldsWithCallat(tags []tag.Tag) []zap.Field {
	_ = "STUB: not implemented"
	return nil
}

func (lg *loggerImpl) buildFields(tags []tag.Tag) []zap.Field {
	_ = "STUB: not implemented"
	return nil
}

// ignore empty field(which can be constructed manually)

func setDefaultMsg(msg string) string { _ = "STUB: not implemented"; return "" }

func (lg *loggerImpl) Debug(msg string, tags ...tag.Tag) { _ = "STUB: not implemented"; return }

func (lg *loggerImpl) Info(msg string, tags ...tag.Tag) { _ = "STUB: not implemented"; return }

func (lg *loggerImpl) Warn(msg string, tags ...tag.Tag) { _ = "STUB: not implemented"; return }

func (lg *loggerImpl) Error(msg string, tags ...tag.Tag) { _ = "STUB: not implemented"; return }

func (lg *loggerImpl) Fatal(msg string, tags ...tag.Tag) { _ = "STUB: not implemented"; return }

func (lg *loggerImpl) WithTags(tags ...tag.Tag) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}
