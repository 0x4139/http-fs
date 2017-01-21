package server

import (
	"github.com/gin-gonic/gin"
	"path"
	"os"
	"log"
	"time"
	"net/http"
	"fmt"
	"strconv"
	"bytes"
)

var savePath string

func init() {
	if savePath = os.Getenv("SAVEPATH"); savePath == "" {
		panic("Please provide the SAVEPATH environment variable ")
	}
	log.Printf("SAVEPATH: %s", apiKey)
}

func Read(c *gin.Context) {
	filename := c.Query("filename")
	if file, err := os.OpenFile(path.Join(savePath, filename), os.O_RDONLY, 0777); err == nil {
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
		http.ServeContent(c.Writer, c.Request, filename, time.Now(), file)
	} else {
		c.String(http.StatusNotFound, "file not found: %s", filename)
	}

}

func ReadAt(c *gin.Context) {
	filename := c.Query("filename")
	var err error
	var size int
	var offset int64
	if size, err = strconv.Atoi(c.Query("size")); err != nil {
		c.String(http.StatusBadRequest, "invalid size: %s", c.Query("size"))
	}
	if offset, err = strconv.ParseInt(c.Query("offset"), 10, 64); err != nil {
		c.String(http.StatusBadRequest, "invalid offset: %s", c.Query("offset"))
	}

	if file, err := os.OpenFile(path.Join(savePath, filename), os.O_RDONLY, 0777); err == nil {

		data := make([]byte, size)
		_, err = file.ReadAt(data, offset)
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
		http.ServeContent(c.Writer, c.Request, filename, time.Now(), bytes.NewReader(data))
	} else {
		c.String(http.StatusNotFound, "file not found: %s", filename)
	}

}