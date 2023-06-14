package asyncpool

import (
	"log"
	"os"
)

type Logger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type DefaultLogger struct {
	logger *log.Logger
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{
		logger: log.New(os.Stdout, "-------***------", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds|log.Llongfile),
	}
}

func (logger *DefaultLogger) Info(format string, args ...interface{}) {
	logger.logger.Printf(format, args...)
}

func (logger *DefaultLogger) Error(format string, args ...interface{}) {
	logger.logger.Printf(format, args...)
}
