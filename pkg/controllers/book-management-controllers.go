package controllers

import (
	"fmt"
	"net/http"
)

// POST
// Create a book
var CreateBookHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("from CreateBookHandler")
}

// GET
// Get all books
var GetAllBooksHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("from GetAllBooksHandler")

}

// GET
// Get a book by its Id
var GetBookByIdHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("from GetBookByIdHandler")
}

// PUT
// Update a book detail
var UpdateBookHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("from UpdateBookHandler")
}

// DELETE
// Delete a book
var DeleteBookHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("from DeleteBookHandler")
}
