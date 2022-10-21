package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// model
type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func main() {
	fmt.Println("initial code")
	app := fiber.New()

	// Db setup
	dsn := "root:MyPass123!mYsql@tcp(localhost:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Lead{})

	// routes

	// Get all leads
	app.Get("/api/lead", func(c *fiber.Ctx) error {
		var leads []Lead
		db.Find(&leads)

		return c.JSON(leads)
	})

	// Get leads by Id
	app.Get("/api/lead/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var lead Lead
		db.Where("ID=?", id).Find(&lead)
		return c.JSON(lead)
	})

	// Create new lead
	app.Post("/api/lead", func(c *fiber.Ctx) error {
		lead := new(Lead)
		if err := c.BodyParser(lead); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		db.Create(&lead)
		return c.JSON(lead)
	})

	// Delete lead by id
	app.Delete("/api/lead/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var lead Lead
		var leads []Lead
		db.Where("ID=?", id).Delete(&lead)
		db.Find(&leads)
		return c.JSON(leads)
	})

	// Start Server
	log.Fatal(app.Listen(":8080"))
}
