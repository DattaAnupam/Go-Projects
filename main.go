package main

import (
	"log"
	"main/pkg/routes"
	"main/pkg/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// declare new router
	r := mux.NewRouter()

	services.Init()

	// Handle routing
	routes.RegisterBookManagementRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
