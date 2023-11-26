// logsFolderPath :=   //E:/WFH/myLog
// project-log/main.go
package main

import (
	"net/http"

	api "log_app/api"
	"log_app/logger"
)

func main() {
	// Specify the absolute path for the logs folder
	logsFolderPath := "E:/WFH/myLogFile/firstLog"

	// Initialize loggers with specific file paths
	logger.InitLoggersWithPaths(logsFolderPath)

	// Define routes
	http.HandleFunc("/login", api.LoginHandler)

	// Start server
	port := ":8080"
	logger.InfoLogger.Printf("Server started on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error starting server: %v\n", err)
	}
}
