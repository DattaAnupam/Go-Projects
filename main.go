package main

import (
	"fmt"
	"main/pkg/config"
	"main/pkg/controllers"
	"main/pkg/models"
	"main/pkg/routes"
	"main/pkg/utils"
)

func main() {
	fmt.Println("basic structure")
	routes.TryFunc()
	controllers.TryControllers()
	models.TryModels()
	config.TryConfig()
	utils.TryUtils()
}
