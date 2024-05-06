package pomodoro

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	logger := log.New(os.Stdout, "", 0)
	logger.SetPrefix("\npomodoro: ")
	return &Logger{
		logger,
	}
}
