package routes

import (
	"github.com/gin-gonic/gin"
	"personal-website/app/middleware"
	"personal-website/app/service/v1"
)

func SetRouter(r *gin.Engine) {
	V1Api := r.Group("api/v1")
	{
		// register a user
		V1Api.POST("users/register", v1.RegisterUser)
		// login a user
		V1Api.POST("users/login", v1.LoginUser)
	}

	V1ApiAuth := r.Group("api/v1")
	V1ApiAuth.Use(middleware.AuthRequired)
	{
		// logout a user
		V1ApiAuth.GET("users/logout", v1.LogoutUser)
	}
}
