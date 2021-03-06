package routes

import (
	"github.com/gin-gonic/gin"
	"personal-website/app/middlewares"
	"personal-website/app/services/v1"
)

func SetRouter(r *gin.Engine) {
	V1Api := r.Group("api/v1")
	{
		// register a user
		V1Api.POST("users/register", v1.RegisterUser)

		// login a user
		V1Api.POST("users/login", v1.LoginUser)

		// list a blog
		V1Api.POST("blogs/list", v1.ListBlogs)
	}

	V1ApiAuth := r.Group("api/v1")
	V1ApiAuth.Use(middlewares.AuthRequired)
	{
		// logout a user
		V1ApiAuth.GET("users/logout", v1.LogoutUser)

		// create a blog
		V1ApiAuth.POST("blogs/create", v1.CreateBlog)

		// delete a blog
		V1ApiAuth.POST("blogs/delete", v1.DeleteBlog)
	}
}
