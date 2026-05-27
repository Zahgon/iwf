package cadence

import (
	"go.uber.org/zap"
)

type loggerImpl struct {
	zlogger *zap.Logger
}

func buildZapFields(keyvals []interface{}) []zap.Field { _ = "STUB: not implemented"; return nil }

func (a *loggerImpl) Debug(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

func (a *loggerImpl) Info(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

func (a *loggerImpl) Warn(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }

func (a *loggerImpl) Error(msg string, keyvals ...interface{}) { _ = "STUB: not implemented"; return }
