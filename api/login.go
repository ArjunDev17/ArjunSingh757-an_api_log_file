// api/login.go
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"log_app/logger"
)

// User represents a user for login
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// Log the start of the request
	logger.InfoLogger.Printf("Request started at %s\n", startTime.Format(time.RFC3339))

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.ErrorLogger.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Simulate authentication logic
	if user.Username == "example" && user.Password == "password" {
		// Log successful login
		logger.InfoLogger.Printf("Login successful for user: %s\n", user.Username)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Login successful")
	} else {
		// Log failed login attempt
		logger.WarningLogger.Printf("Login failed for user: %s\n", user.Username)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}

	// Log the end of the request and the total time taken
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	logger.InfoLogger.Printf("Request completed at %s. Duration: %v\n", endTime.Format(time.RFC3339), duration)
}
