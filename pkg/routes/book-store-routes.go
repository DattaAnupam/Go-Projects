package routes

import (
	"main/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRouters = func(router *mux.Router) {
	router.HandleFunc("/api/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/book/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
