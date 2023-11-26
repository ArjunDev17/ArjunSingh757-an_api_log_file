// main.go
package main

import (
	"net/http"
	"os"

	"log_app/api"
	"log_app/logger"
)

func main() {
	// Initialize loggers
	infoFile, _ := os.Create("info.log")
	warningFile, _ := os.Create("warning.log")
	errorFile, _ := os.Create("error.log")

	logger.InitLoggers(infoFile, warningFile, errorFile)

	// Define routes
	http.HandleFunc("/login", api.LoginHandler)

	// Start server
	port := ":8088"
	logger.InfoLogger.Printf("Server started on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error starting server: %v\n", err)
	}
}
