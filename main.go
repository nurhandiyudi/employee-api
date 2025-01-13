package main

import (
	"log"
	//"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"employee-api/router"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	
	// Register Routes
	router.RegisterRoutes(e)

	// Start the server
	log.Println("Starting server on :3000")
	if err := e.Start(":3000"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
