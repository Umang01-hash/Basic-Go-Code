package main

import (
	"fmt"
	"net/http"
)

// eventHandler defines the functionality for the handler
func eventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		w.Write([]byte("invalid method")) // [] - notation for slice data set in golang
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to GoFr's event!"))
}

func main() {
	// Register handler function for `/event` in default ServeMux
	http.HandleFunc("/event", eventHandler)

	// Inform the user that the server is running
	fmt.Println("Server is running on http://localhost:8020")

	// Start the HTTP server
	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		fmt.Printf("Server encountered an error: %s\n", err)
	}
}
