package controllers

import (
	"github.com/anupam/crm-with-sql-fiber/pkg/config"
	"github.com/anupam/crm-with-sql-fiber/pkg/model"
	"github.com/gofiber/fiber/v2"
)

// Get all leads
func GetLeads(c *fiber.Ctx) {
	db := config.GetDB()
	var leads []model.Lead

	db.Find(&leads)
	c.JSON(leads)
}

// Get a single lead
var GetLead = func(c *fiber.Ctx) {
	// Get Id
	id := c.Params("id")
	db := config.GetDB()

	var lead model.Lead
	db.Find(&lead, id)
	// db.Where("ID=?", id).Find(&lead)

	c.JSON(lead)
}

// Create a new Lead
var CreateLead = func(c *fiber.Ctx) {
	db := config.GetDB()
	lead := new(model.Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).SendString(err.Error())
		return
	}

	db.Create(&lead)
	c.JSON(lead)
}

// Delete an existing Lead
var DeleteLead = func(c *fiber.Ctx) {
	db := config.GetDB()
	id := c.Params("id")

	var lead model.Lead
	db.First(&lead, id)

	if lead.Name == "" {
		c.Status(500).SendString("No lead found with ID")
	}

	c.JSON("Id deleted successfully")
}
