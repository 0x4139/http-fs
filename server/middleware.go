package server

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"net/http"
)

var apiKey string

func init() {
	if apiKey = os.Getenv("KEY"); apiKey == "" {
		panic("Please provide the KEY environment variable ")
	}
	log.Printf("KEY: %s", apiKey)
}

func ApiKeyMiddleWare(c *gin.Context) {
	if c.Query("key") != apiKey {
		c.String(http.StatusForbidden, "Invalid key %s", c.Query("key"))
		c.Abort()
	}else{
		c.Next()
	}
}
