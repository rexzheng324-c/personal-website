package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//默认为监听8080端口
	r.Run(":8000")
}
