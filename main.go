package main

import (
	"fmt"
	"log"

	"github.com/anupam/crm-with-go-fiber-mysql/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("initial code")
	app := fiber.New()

	// routes
	routes.SetupRoutes(app)

	// Start Server
	log.Fatal(app.Listen(":8080"))
}
