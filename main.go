package main

import (
	"log"
	"main/pkg/routes"
	"main/pkg/services"
	"main/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	routes.TryFunc()
	utils.TryUtils()
	services.TryServices()

	// declare new router
	r := mux.NewRouter()
	// Handle routing
	routes.RegisterBookManagementRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
