package services

import (
	"fmt"
	"main/pkg/models"

	"gorm.io/gorm"
)

func TryServices() {
	fmt.Println("from main/pkg/services/book-management-services.go")
}

var db *gorm.DB

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
