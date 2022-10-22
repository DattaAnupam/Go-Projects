package routes

import (
	"github.com/anupam/crm-with-go-fiber-mysql/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

var SetupRoutes = func(app *fiber.App) {
	// setup database
	controllers.SetupDb()

	// Get all leads
	app.Get("/api/lead", controllers.GetAllLeads)

	// Get leads by Id
	app.Get("/api/lead/:id", controllers.GetLeadByID)

	// Create new lead
	app.Post("/api/lead", controllers.CreateLead)

	// Update lead
	app.Put("/api/lead/:id", controllers.UpdateLead)

	// Delete lead by id
	app.Delete("/api/lead/:id", controllers.DeleteLeadByID)
}
