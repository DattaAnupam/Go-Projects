package main

import (
	"fmt"
	"log"

	"github.com/anupam/crm-with-go-fiber-mysql/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("initial code")
	app := fiber.New()

	// routes

	// Get all leads
	app.Get("/api/lead", controllers.GetAllLeads)

	// Get leads by Id
	app.Get("/api/lead/:id", controllers.GetLeadByID)

	// Create new lead
	app.Post("/api/lead", controllers.CreateLead)

	// Delete lead by id
	app.Delete("/api/lead/:id", controllers.DeleteLeadByID)

	// Start Server
	log.Fatal(app.Listen(":8080"))
}
