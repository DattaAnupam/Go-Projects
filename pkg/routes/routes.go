package routes

import (
	"github.com/anupam/crm-with-sql-fiber/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

var SetupRoute = func(app *fiber.App) {
	app.Get("/api/lead", controllers.GetLeads)
	app.Get("/api/lead/:id")
	app.Post("/api/lead")
	app.Delete("/api/lead/:id")
}
