package main

import (
	"net/http"
	"os"

	"log_app/api"
	"log_app/logger"
)

func main() {
	// Create a folder for logs
	err := os.MkdirAll("logger/logs", os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Initialize loggers
	infoFile, _ := os.Create("logger/logs/info.log")
	warningFile, _ := os.Create("logger/logs/warning.log")
	errorFile, _ := os.Create("logger/logs/error.log")

	// Initialize the logger with the log file handles
	logger.InitLoggers(infoFile, warningFile, errorFile)

	// Define routes
	http.HandleFunc("/login", api.LoginHandler)

	// Start server
	port := ":8080"
	logger.InfoLogger.Printf("Server started on port %s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error starting server: %v\n", err)
	}
}
