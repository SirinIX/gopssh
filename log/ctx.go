package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type CtxLogger struct {
	logger *logrus.Entry
	Fields map[string]interface{}
}

func NewCtxLogger(fields map[string]interface{}) *CtxLogger {
	return &CtxLogger{
		logger: logrus.WithFields(fields),
		Fields: fields,
	}
}

func (l *CtxLogger) Debug(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}

func (l *CtxLogger) Info(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

func (l *CtxLogger) Warn(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *CtxLogger) Error(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}
