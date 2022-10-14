package main

import (
	"log"
	"main/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// declare new router
	r := mux.NewRouter()
	// Handle routing
	routes.RegisterBookManagementRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
