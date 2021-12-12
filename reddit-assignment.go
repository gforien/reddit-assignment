package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()
	counter := 0

	// Get OK
	r.GET("/ok", func(c *gin.Context) {
		c.String(http.StatusOK, "gin OK")
	})

	// Get counter
	r.GET("/count", func(c *gin.Context) {
		c.String(http.StatusOK, "%d", counter)
	})

	// Increment counter
	r.POST("/inc", func(c *gin.Context) {
		counter++
		c.String(http.StatusOK, "%d", counter)
	})

	// Decrement counter
	r.POST("/dec", func(c *gin.Context) {
		counter--
		c.String(http.StatusOK, "%d", counter)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":5000")
}
