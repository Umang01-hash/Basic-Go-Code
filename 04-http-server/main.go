package main

import (
	"fmt"
	"net/http"
)

// eventHandler defines the functionality for the handler
func eventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to GoFr's event!"))
}

func main() {
	// Register handler function for `/event` in default ServeMux
	http.HandleFunc("/event", eventHandler)

	// Define the HTTP port
	port := 8080
	serverAddress := fmt.Sprintf(":%d", port)

	// Inform the user that the server is running
	fmt.Printf("Server is running on http://localhost%s\n", serverAddress)

	// Start the HTTP server
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		fmt.Printf("Server encountered an error: %s\n", err)
	}
}
