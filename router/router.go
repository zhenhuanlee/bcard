package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"message": "page not found"})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome",
		})
	})

	r.RunTLS(":8080", "./config/server", "./config/server.key")

	// apiV1 := r.Group("/api/v1")
	// apiV1.GET("/tags", v1.GetTags)
}
