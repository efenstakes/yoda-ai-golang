package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// create default server
	r := gin.Default()

	// index route
	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// start server
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
