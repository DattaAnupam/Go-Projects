package main

import (
	"fmt"

	"github.com/anupam/go-fiber-crm-basic/database"
	"github.com/anupam/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/lead", lead.GetLeads)
	app.Get("/api/lead/:id", lead.GetLead)
	app.Post("/api/lead", lead.NewLead)
	app.Delete("/api/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	dsn := "root:MyPass123!mYsql@tcp(localhost:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

	mySqlDb, err := database.DBConn.DB()
	if err != nil {
		panic("Error in getting db")
	}

	defer mySqlDb.Close()
}
