package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookName string `json:"bookname"`
	Author   string `json:"author"`
	Isbn     string `json:"isbn"`
}

func TryModels() {
	fmt.Println("from pkg/models/book-management-models.go")
}
