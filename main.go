package main

import (
	"fmt"
	"main/pkg/controllers"
	"main/pkg/routes"
)

func main() {
	fmt.Println("basic structure")
	routes.TryFunc()
	controllers.TryControllers()
}
