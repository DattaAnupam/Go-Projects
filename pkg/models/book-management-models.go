package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	BookName string `json:"bookname"`
	Author   string `json:"author"`
	Isbn     string `json:"isbn"`
}
