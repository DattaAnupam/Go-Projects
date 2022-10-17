package routes

import (
	"main/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookManagementRoutes = func(router *mux.Router) {
	// Get all books - route
	router.HandleFunc("/api/books", controllers.GetAllBooksHandler).Methods("GET")
	// Get a book - route
	router.HandleFunc("/api/books/{bookId}", controllers.GetBookByIdHandler).Methods("GET")
	// Create book - route
	router.HandleFunc("/api/books", controllers.CreateBookHandler).Methods("POST")
	// Update book details - route
	router.HandleFunc("/api/books/{bookId}", controllers.UpdateBookHandler).Methods("PUT")
	// Delete book - route
	router.HandleFunc("/api/books/{bookId}", controllers.DeleteBookHandler).Methods("DELETE")
}
