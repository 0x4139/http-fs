package server

import (
	"github.com/gin-gonic/gin"
	"path"
	"os"
	"log"
)

var savePath string

func init() {
	if savePath = os.Getenv("SAVEPATH"); savePath == "" {
		panic("Please provide the SAVEPATH environment variable ")
	}
	log.Printf("SAVEPATH: %s", apiKey)
}

func Read(c *gin.Context) {
	c.File(path.Join(savePath, c.Query("filename")))
}