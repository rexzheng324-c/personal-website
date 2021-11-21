package main

import (
	"personal-website/database/model"
	"personal-website/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
