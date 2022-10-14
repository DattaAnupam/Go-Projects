package main

import (
	"fmt"
	"log"
	"main/pkg/config"
	"main/pkg/models"
	"main/pkg/routes"
	"main/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("basic structure")
	routes.TryFunc()
	models.TryModels()
	config.TryConfig()
	utils.TryUtils()

	// declare new router
	r := mux.NewRouter()
	// Handle routing
	// todo

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
