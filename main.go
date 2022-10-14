package main

import (
	"fmt"
	"log"
	"main/pkg/config"
	"main/pkg/controllers"
	"main/pkg/models"
	"main/pkg/routes"
	"main/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("basic structure")
	routes.TryFunc()
	controllers.TryControllers()
	models.TryModels()
	config.TryConfig()
	utils.TryUtils()

	r := mux.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
