package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetOutput(os.Stdout)
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"location": "root",
		})
	})
	r.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"location": "list",
		})
	})
	r.GET("/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"location": "add",
		})
	})
	r.GET("/remove", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"location": "remove",
		})
	})
	return r
}
