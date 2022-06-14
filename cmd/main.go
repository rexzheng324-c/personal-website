package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"personal-website/app/databases/mysql"
	"personal-website/app/routes"
)

func main() {
	r := gin.Default()

	// store session
	r.Use(sessions.Sessions("SESSIONID", cookie.NewStore([]byte("secret"))))

	// set the router
	routes.SetRouter(r)

	// init the mysql client
	_, err := mysql.GetDb("config/local_config.ini")
	if err != nil {
		panic(err)
	}

	err = r.Run(":8000")
	if err != nil {
		panic(err)
	}
}
