// project-log/logger/logger.go
package logger

import (
	"log"
	"os"
	"path/filepath"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func InitLoggers(infoHandle, warningHandle, errorHandle *os.File) {
	// Initialize loggers with the log file handles
	InfoLogger = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitLoggersWithPaths(logsFolderPath string) {
	// Create the logs folder if it doesn't exist
	err := os.MkdirAll(logsFolderPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Create log file paths using the specified logs folder
	infoLogPath := filepath.Join(logsFolderPath, "info.log")
	warningLogPath := filepath.Join(logsFolderPath, "warning.log")
	errorLogPath := filepath.Join(logsFolderPath, "error.log")

	// Create or open log files
	infoFile, err := os.Create(infoLogPath)
	if err != nil {
		panic(err)
	}

	warningFile, err := os.Create(warningLogPath)
	if err != nil {
		panic(err)
	}

	errorFile, err := os.Create(errorLogPath)
	if err != nil {
		panic(err)
	}

	// Initialize loggers with the log file handles
	InitLoggers(infoFile, warningFile, errorFile)
}
