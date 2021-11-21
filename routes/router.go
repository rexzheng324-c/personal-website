package routes

import (
	"github.com/gin-gonic/gin"
	v1 "personal-website/service/v1"
)

func InitRouter() {
	r := gin.New()
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	{
		// user api
		auth.GET("users/:id", v1.GetUser)
		auth.GET("users", v1.ListUsers)
		auth.POST("users", v1.CreateUser)
		auth.PUT("users/:id", v1.UpdateUser)
		auth.DELETE("users/:id", v1.DeleteUser)
	}

	_ = r.Run(":8000")

}
