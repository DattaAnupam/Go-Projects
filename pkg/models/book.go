package models

import (
	"main/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Create a new book record inside db
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// Get all books from db
func GetAllBooks() []Book {
	var books []Book

	db.Find(&books)

	return books
}

// Get a book from db
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book

	db.Where("ID=?", Id).Find(&getBook)

	return &getBook, db
}

// Delete a book from db
func DeleteBookById(Id int64) Book {
	var deletedBook Book

	db.Where("ID=?", Id).Delete(&deletedBook)

	return deletedBook
}
