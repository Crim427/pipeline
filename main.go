package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetOutput(os.Stdout)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"location": "root",
		})
	})
	r.GET("/list", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"location": "list",
		})
	})
	r.GET("/add", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"location": "add",
		})
	})
	r.GET("/remove", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"location": "remove",
		})
	})
	r.Run()
}
