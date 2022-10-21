package controllers

import (
	"github.com/anupam/crm-with-sql-fiber/pkg/config"
	"github.com/anupam/crm-with-sql-fiber/pkg/model"
	"github.com/gofiber/fiber/v2"
)

// Get all leads
func GetLeads(c *fiber.Ctx) error {
	db := config.GetDB()
	var leads []model.Lead

	db.Find(&leads)
	return c.JSON(leads)
}

// Get a single lead
var GetLead = func(c *fiber.Ctx) error {
	// Get Id
	id := c.Params("id")
	db := config.GetDB()

	var lead model.Lead
	db.Find(&lead, id)
	// db.Where("ID=?", id).Find(&lead)

	return c.JSON(lead)
}

// Create a new Lead
var CreateLead = func(c *fiber.Ctx) error {
	db := config.GetDB()
	lead := new(model.Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&lead)
	return c.JSON(lead)
}

// Delete an existing Lead
var DeleteLead = func(c *fiber.Ctx) error {
	db := config.GetDB()
	id := c.Params("id")

	var lead model.Lead
	db.First(&lead, id)

	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with ID")
	}

	return c.JSON("Id deleted successfully")
}
