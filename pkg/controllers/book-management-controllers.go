package controllers

import (
	"encoding/json"
	"fmt"
	"main/pkg/services"
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
	books := services.GetAllBooks()

	// Create response
	res, _ := json.Marshal(books)

	// Set Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
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
