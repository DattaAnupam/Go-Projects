package main

import (
	"encoding/json"
	"log"
	"main/pkg/routes"
	"main/pkg/services"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// GET
// Get all books
func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

// GET
// Get a book by its Id
func getBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get params
	params := mux.Vars(r)

	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// Return Empty book if not found
	json.NewEncoder(w).Encode(&Book{})
}

// POST
// Create a book
func createBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	// Decode request json body
	_ = json.NewDecoder(r.Body).Decode(&book)
	// Assign random Id to the new book
	book.ID = strconv.Itoa(rand.Intn(1000000))
	// Add new book to the books slice
	books = append(books, book)

	// return newly create book
	json.NewEncoder(w).Encode(book)
}

// PUT
// Update a book by its Id
func updateBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	// Decode request json body
	_ = json.NewDecoder(r.Body).Decode(&book)
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			// remove the book from the list
			books = append(books[:index], books[index+1:]...)
			// add new book
			book.ID = params["id"]
			books = append(books, book)
			break
		}
	}

	// return newly create book
	json.NewEncoder(w).Encode(book)
}

// DELETE
// Delete a book by its Id
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	// return all the books
	json.NewEncoder(w).Encode(books)
}

func main() {
	// declare new router
	r := mux.NewRouter()

	services.Init()

	// Handle routing
	routes.RegisterBookManagementRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
