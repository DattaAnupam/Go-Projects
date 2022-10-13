package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Connect to the Sql Db
func Connect() {
	// dsn = Data Source Name
	dsn := "root:<MySqlPassword>@tcp(localhost:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

// return db
func GetDB() *gorm.DB {
	return db
}
