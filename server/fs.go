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
	"io"
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
		defer file.Close()
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
		defer file.Close()
		data := make([]byte, size)
		_, err = file.ReadAt(data, offset)
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
		http.ServeContent(c.Writer, c.Request, filename, time.Now(), bytes.NewReader(data))
	} else {
		c.String(http.StatusNotFound, "file not found: %s", filename)
	}

}

func Write(c *gin.Context) {
	filename := c.Query("filename")
	if file, _, err := c.Request.FormFile("data"); err != nil {
		c.String(http.StatusBadRequest, "error while writing file %s", err)
	} else {
		defer file.Close()
		if filetoSave, err := os.OpenFile(path.Join(path.Join(savePath, filename)), os.O_WRONLY | os.O_CREATE, 0777); err != nil {
			c.String(http.StatusBadRequest, "error while writing file %s", err)
		} else {
			defer file.Close()
			if _,err:=io.Copy(filetoSave,file);err!=nil{
				c.String(http.StatusBadRequest, "error while writing file %s", err)
			}
		}
	}
}