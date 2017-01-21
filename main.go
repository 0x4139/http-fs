package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/0x4139/http-fs/server"
)

func main() {

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(server.ApiKeyMiddleWare)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/read", server.Read)
	r.GET("/readat",server.ReadAt)
	r.Run() // listen and serve on 0.0.0.0:8080
}