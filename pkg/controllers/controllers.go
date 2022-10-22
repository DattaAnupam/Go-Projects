package controllers

import (
	"github.com/anupam/crm-with-go-fiber-mysql/pkg/config"
	"github.com/anupam/crm-with-go-fiber-mysql/pkg/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

var SetupDb = func() {
	db = config.InitDB()
}

// Get all leads
var GetAllLeads = func(c *fiber.Ctx) error {
	var leads []model.Lead
	db.Find(&leads)

	return c.JSON(leads)
}

// Get a single lead by ID
var GetLeadByID = func(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead model.Lead
	db.Where("ID=?", id).Find(&lead)
	return c.JSON(lead)
}

// Create a New Lead
var CreateLead = func(c *fiber.Ctx) error {
	lead := new(model.Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&lead)
	return c.JSON(lead)
}

// Update a Lead
var UpdateLead = func(c *fiber.Ctx) error {
	// Save Incoming Request Info
	newLeadInfo := new(model.Lead)
	if err := c.BodyParser(newLeadInfo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	// Find Particular Lead by Id
	id := c.Params("ID")
	var oldLead model.Lead
	db.Model(model.Lead{}).Where("ID=?", id).Updates(newLeadInfo)
	db.Where("ID=?", id).Find(&oldLead)

	return c.JSON(oldLead)
}

// Delete a Lead by ID
var DeleteLeadByID = func(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead model.Lead
	var leads []model.Lead
	db.Where("ID=?", id).Delete(&lead)
	db.Find(&leads)
	return c.JSON(leads)
}
