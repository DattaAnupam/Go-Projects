package config

import (
	"log"

	"github.com/anupam/crm-with-sql-fiber/pkg/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// Connect to Mysql database
var Connect = func() {
	dsn := "root:MyPass123!mYsql@tcp(localhost:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DBConn = db
}

// Get database ref
func GetDB() *gorm.DB {
	return DBConn
}

// Initiate Database
var InitDB = func() {
	// connect to DB
	Connect()

	DBConn.AutoMigrate(&model.Lead{})
}

// Initiate Server
var InitServer = func(app *fiber.App) {
	log.Fatal(app.Listen(":8080"))
}
