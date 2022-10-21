package main

import (
	"fmt"
	"log"

	"github.com/anupam/crm-with-go-fiber-mysql/pkg/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("initial code")
	app := fiber.New()

	// Db setup
	dsn := "root:MyPass123!mYsql@tcp(localhost:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.Lead{})

	// routes

	// Get all leads
	app.Get("/api/lead", func(c *fiber.Ctx) error {
		var leads []model.Lead
		db.Find(&leads)

		return c.JSON(leads)
	})

	// Get leads by Id
	app.Get("/api/lead/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var lead model.Lead
		db.Where("ID=?", id).Find(&lead)
		return c.JSON(lead)
	})

	// Create new lead
	app.Post("/api/lead", func(c *fiber.Ctx) error {
		lead := new(model.Lead)
		if err := c.BodyParser(lead); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		db.Create(&lead)
		return c.JSON(lead)
	})

	// Delete lead by id
	app.Delete("/api/lead/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var lead model.Lead
		var leads []model.Lead
		db.Where("ID=?", id).Delete(&lead)
		db.Find(&leads)
		return c.JSON(leads)
	})

	// Start Server
	log.Fatal(app.Listen(":8080"))
}
