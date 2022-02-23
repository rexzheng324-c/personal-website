package main

import (
	"github.com/gin-gonic/gin"
	"personal-website/routes"
)

func main() {
	r := gin.Default()

	routes.SetRouter(r)

	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
}
