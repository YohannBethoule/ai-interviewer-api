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
	r.GET("/prompt", testPrompt())
	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})
	r.Run()
}
