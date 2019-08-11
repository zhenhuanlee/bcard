package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": 123,
		})
	})

	r.RunTLS(":8080", "./config/server.pem", "./config/server.key")
}
