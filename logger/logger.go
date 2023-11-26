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
	// Get the absolute path of the logs folder
	logsFolder, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), "logs"))
	if err != nil {
		panic(err)
	}

	// Create the logs folder if it doesn't exist
	err = os.MkdirAll(logsFolder, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Create log file paths
	infoLogPath := filepath.Join(logsFolder, "info.log")
	warningLogPath := filepath.Join(logsFolder, "warning.log")
	errorLogPath := filepath.Join(logsFolder, "error.log")

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
	InfoLogger = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(warningFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
