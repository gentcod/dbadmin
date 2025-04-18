package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:    log.New(os.Stdout, "\nINFO: ", log.Ltime),
		warningLogger: log.New(os.Stdout, "\nWARNING: ", log.Ltime),
		errorLogger:   log.New(os.Stdout, "\nERROR: ", log.Ltime),
	}
}

func (logger *Logger) Info(value string) {
	logger.infoLogger.Print(value)
}

func (logger *Logger) Warn(value string) {
	logger.warningLogger.Print(value)
}

func (logger *Logger) Error(err error) {
	value := fmt.Sprintf("%v", err)
	logger.errorLogger.Print(value)
}
