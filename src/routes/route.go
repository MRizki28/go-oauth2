package routes

import (
	"github.com/MRizki28/go-oauth2/src/config"
	"github.com/MRizki28/go-oauth2/src/service"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	config.Oauth()
	r := gin.Default()
	r.GET("/", HelloWorld)
	r.GET("/login", service.Login)
	r.GET("auth/callback", service.HandleCallback)
	return r
}

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}