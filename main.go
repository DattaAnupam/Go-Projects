package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
var movies []Movie

// GET
// Get all movies
func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

// GET
// Get a movie by id
func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	// movie found status flag
	foundMovie := false

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			foundMovie = true
			break
		}
	}

	// If not found return empty data
	if !foundMovie {
		json.NewEncoder(w).Encode(Movie{})
	}
}

// POST
// Create a movie
func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie Movie

	w.Header().Set("Content-Type", "application/json")

	// Decode request
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Set Id of new movie
	movie.ID = strconv.Itoa(rand.Intn(10000))

	// Add new movie to the collection
	movies = append(movies, movie)

	// Return newly created movie
	json.NewEncoder(w).Encode(movie)
}

// PUT
// Update a movie by id
func updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var newMovie Movie

	_ = json.NewDecoder(r.Body).Decode(&newMovie)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			newMovie.ID = params["id"]
			movies[index] = newMovie
			// Return updated movie
			json.NewEncoder(w).Encode(newMovie)

			break
		}
	}
}

// DELETE
// Delete a movie by id
func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// return rest of the movies
	json.NewEncoder(w).Encode(movies)
}

func main() {
	// Create new Router
	r := mux.NewRouter()

	// Init movies
	movies = append(movies, Movie{ID: "1", Isbn: "123456", Title: "First Movie", Director: &Director{Firstname: "Anupam", Lastname: "Datta"}})
	movies = append(movies, Movie{ID: "2", Isbn: "345678", Title: "Second Movie", Director: &Director{Firstname: "Second", Lastname: "Director"}})

	// Create Routes and Handlers
	r.HandleFunc("/api/movies", getMoviesHandler).Methods("GET")
	r.HandleFunc("/api/movies/{id}", getMovieHandler).Methods("GET")
	r.HandleFunc("/api/movies", createMovieHandler).Methods("POST")
	r.HandleFunc("/api/movies/{id}", updateMovieHandler).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", deleteMovieHandler).Methods("DELETE")

	// Initiate and Start Server
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
