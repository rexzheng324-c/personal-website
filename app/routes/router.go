package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
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
		// logout a user
		V1Api.GET("users/logout", v1.LogoutUser)
	}

	V1ApiAuth := r.Group("api/v1")
	V1ApiAuth.Use(middleware.AuthRequired)
	{
		V1ApiAuth.GET("users/me", me)
	}
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(middleware.UserKey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
