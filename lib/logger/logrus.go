package logger

import (
	"github.com/sirupsen/logrus"
)

// logrusLogger wrap logrus.Logger
type logrusLogger struct {
	logger *logrus.Logger
}

// NewLogrusLogger create logrusLogger
func NewLogrusLogger() *logrusLogger {
	return &logrusLogger{
		logger: logrus.New(),
	}
}

// Info
func (l *logrusLogger) Info(msg string) {
	l.logger.Info(msg)
}
