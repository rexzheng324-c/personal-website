package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var UserKey = "user"
var UserIdKey = "userId"

// AuthRequired is a simple middlewares to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get(UserKey)
	if userId == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Set(UserIdKey, userId)
	// Continue down the chain to handler etc
	c.Next()
}
