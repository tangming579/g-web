package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.Println("start")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
