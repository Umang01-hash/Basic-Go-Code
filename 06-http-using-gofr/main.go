package main

import (
	"gofr.dev/pkg/gofr"
)

// eventHandler defines the functionality for the handler
func eventHandler(ctx *gofr.Context) (interface{}, error) {
	return "Welcome to GoFr's event!", nil
}

func main() {
	app := gofr.New()

	app.GET("/event", eventHandler)

	app.Start()
}
