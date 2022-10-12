package config

import (
	"gorm.io/driver/sqlite"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Connect to the Sql Db
func Connect() {
	// dsn = Data Source Name
	// dsn := "user:pass@tcp(localhost:8080)/simpledb?charset=utf8mb4&parseTime=True&loc=Local"
	// d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

// return db
func GetDB() *gorm.DB {
	return db
}
