package common

import (
	"strings"

	"github.com/sirupsen/logrus"
)

//Creates a new logger instance
func NewLogger(loglevel string) logrus.FieldLogger {
	logger := logrus.New()
	if level, err := logrus.ParseLevel(strings.ToLower(loglevel)); err == nil {
		logger.Level = level
	}
	return logger
}
