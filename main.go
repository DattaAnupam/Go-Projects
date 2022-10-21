package main

import (
	"fmt"
	"log"

	"github.com/anupam/crm-with-go-fiber-mysql/pkg/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Db Setup
var InitDB = func() *gorm.DB {
	dsn := "root:MyPass123!mYsql@tcp(localhost:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	d.AutoMigrate(&model.Lead{})
	return d
}

var db = InitDB()

// Controller Functions
var GetAllLeads = func(c *fiber.Ctx) error {
	var leads []model.Lead
	db.Find(&leads)

	return c.JSON(leads)
}

var GetLeadByID = func(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead model.Lead
	db.Where("ID=?", id).Find(&lead)
	return c.JSON(lead)
}

var CreateLead = func(c *fiber.Ctx) error {
	lead := new(model.Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&lead)
	return c.JSON(lead)
}

var DeleteLeadByID = func(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead model.Lead
	var leads []model.Lead
	db.Where("ID=?", id).Delete(&lead)
	db.Find(&leads)
	return c.JSON(leads)
}

func main() {
	fmt.Println("initial code")
	app := fiber.New()

	// Db setup
	// dsn := "root:MyPass123!mYsql@tcp(localhost:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// db.AutoMigrate(&model.Lead{})

	// routes

	// Get all leads
	app.Get("/api/lead", GetAllLeads)

	// Get leads by Id
	app.Get("/api/lead/:id", GetLeadByID)

	// Create new lead
	app.Post("/api/lead", CreateLead)

	// Delete lead by id
	app.Delete("/api/lead/:id", DeleteLeadByID)

	// Start Server
	log.Fatal(app.Listen(":8080"))
}
