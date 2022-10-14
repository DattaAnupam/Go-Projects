package services

import (
	"main/pkg/config"
	"main/pkg/models"

	"gorm.io/gorm"
)

var db *gorm.DB

var Init = func() {
	// Connect to DB
	config.Connect()
	// set the DB
	db = config.GetDB()

	db.AutoMigrate(&models.Book{})
}

// Create new book record inside DB
func CreateBook(b *models.Book) *models.Book {
	db.Create(&b)
	return b
}

// Get all books from db
func GetAllBooks() []models.Book {
	var books []models.Book
	db.Find(&books)
	return books
}

// Get a book by its ID
func GetBookById(Id int64) *models.Book {
	var book models.Book
	db.Where("ID=?", Id).Find(&book)
	return &book
}

// Delete book
func DeleteBook(Id int64) *models.Book {
	var book models.Book
	db.Where("ID=?", Id).Delete(&book)
	return &book
}
