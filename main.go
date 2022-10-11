package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie Model
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director Model
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Slice of movies, dummy for db
var Movies []Movie

// GET
// Get all movies
func getMoviesHandler(w http.ResponseWriter, r *http.Request) {

}

// GET
// Get a movie by id
func getMovieHandler(w http.ResponseWriter, r *http.Request) {

}

// POST
// Create a movie
func createMovieHandler(w http.ResponseWriter, r *http.Request) {

}

// PUT
// Update a movie by id
func updateMovieHandler(w http.ResponseWriter, r *http.Request) {

}

// DELETE
// Delete a movie by id
func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Create new Router
	r := mux.NewRouter()

	// Create Routes and Handlers
	r.HandleFunc("api/movies", getMoviesHandler).Methods("GET")
	r.HandleFunc("api/movies/{id}", getMovieHandler).Methods("GET")
	r.HandleFunc("api/movies", createMovieHandler).Methods("POST")
	r.HandleFunc("api/movies/{id}", updateMovieHandler).Methods("PUT")
	r.HandleFunc("api/movies/{id}", deleteMovieHandler).Methods("DELETE")

	// Initiate and Start Server
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
