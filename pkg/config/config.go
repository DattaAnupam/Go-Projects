package config

import (
	"github.com/anupam/crm-with-go-fiber-mysql/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Db Setup
var InitDB = func() *gorm.DB {
	// Change dsn username, password, port number (if need) and database name
	// e.g.
	// username: root, password: pass123, port on which mysql is running: 3306, dbname: mytestdb
	dsn := "root:pass123@tcp(localhost:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	d.AutoMigrate(&model.Lead{})
	return d
}
