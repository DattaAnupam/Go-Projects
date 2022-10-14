package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

/*
Connect to MySql DB

Database Name: mytestdb
Username: root
*/
var Connect = func() {
	// dsn = data source name
	dsn := "root:MyPass123!mYsql@tcp(localhost:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db = d
}

// return db
var GetDB = func() *gorm.DB {
	return db
}
