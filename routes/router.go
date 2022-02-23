package routes

import (
	"github.com/gin-gonic/gin"
	v1 "personal-website/service/v1"
)

func SetRouter(r *gin.Engine) {
	V1Api := r.Group("api/v1")
	{
		// health check
		V1Api.GET("health/check")
		// user api by auth
		V1Api.GET("users/:id", v1.GetUser)
		V1Api.GET("users", v1.ListUsers)
		V1Api.POST("users", v1.CreateUser)
		V1Api.PUT("users/:id", v1.UpdateUser)
	}
}
