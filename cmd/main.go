package main

import (
	"github.com/gin-gonic/gin"
	"personal-website/database/model"
)

func main() {
	r := gin.Default()

	model.InitDb()
	//默认为监听8080端口
	r.Run(":8000")
}
