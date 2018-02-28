package main

import "github.com/qclaogui/goforum/routes"

func main() {

	forum := routes.InitRoutes()
	forum.Run(":8321")
}
