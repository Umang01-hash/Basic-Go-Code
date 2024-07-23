package main

import (
	"gofr.dev/pkg/gofr"
)

// eventHandler defines the functionality for the handler
func eventHandler(ctx *gofr.Context) (interface{}, error) {
	return "Welcome to GoFr's event!", nil
}

func main() {
	// Creating GoFr's instance
	app := gofr.New()

	// Get endpoint using GoFr
	app.GET("/event", eventHandler)

	// Starting the server
	app.Run()
}
