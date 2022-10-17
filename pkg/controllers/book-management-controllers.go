package controllers

import (
	"encoding/json"
	"fmt"
	"main/pkg/models"
	"main/pkg/services"
	"main/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// POST
// Create a book
var CreateBookHandler = func(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	// Parse new book request
	utils.ParseReqBody(r, book)

	createBook := services.CreateBook(book)

	// Create Response
	res, _ := json.Marshal(createBook)

	// Set Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write(res)
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
	// get query parameters from request
	params := mux.Vars(r)

	// Get book id
	id, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing Book id.")
	}

	book, _ := services.GetBookById(id)

	// Create response
	res, _ := json.Marshal(book)

	// Set Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)

	w.Write(res)
}

// PUT
// Update a book detail
var UpdateBookHandler = func(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	// parse reqest body
	utils.ParseReqBody(r, book)

	// get query parameters from request
	params := mux.Vars(r)

	// get book id and convert it to int
	id, err := strconv.ParseInt(params["bookId"], 0, 0)

	// check for error during parsing
	if err != nil {
		fmt.Println("Error while Parsing Book id.")
	}

	// get the book with id from db
	oldBook, db := services.GetBookById(id)

	// Update book name if given in the input
	if book.BookName != "" {
		oldBook.BookName = book.BookName
	}

	// Update author name if given in the input
	if book.Author != "" {
		oldBook.Author = book.Author
	}

	// Update isbn if given in the input
	if book.Isbn != "" {
		oldBook.Isbn = book.Isbn
	}

	// Save updated details inside db
	db.Save(&oldBook)

	// Create Response
	res, _ := json.Marshal(oldBook)

	// Set Header
	w.Header().Set("Content-Type", "applicaton/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

// DELETE
// Delete a book
var DeleteBookHandler = func(w http.ResponseWriter, r *http.Request) {
	// get query parameters from request
	params := mux.Vars(r)

	// get book id and convert it to int
	id, err := strconv.ParseInt(params["bookId"], 0, 0)

	// check no error during parsing
	if err != nil {
		fmt.Println("Error while Parsing Book id.")
	}

	book := services.DeleteBook(id)

	// Create Response
	res, _ := json.Marshal(book)

	// Set Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}
