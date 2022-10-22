package config

import (
	"github.com/anupam/crm-with-go-fiber-mysql/pkg/model"

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
