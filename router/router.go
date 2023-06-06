package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", testPrompt())
	r.GET("/start", startConversation())
	r.Run()
}
