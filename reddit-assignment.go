package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var db = make(map[string]string)

func setupHTTPServer() *gin.Engine {
	app := gin.Default()
	counter := 0

	// Get OK
	app.GET("/ok", func(c *gin.Context) {
		c.String(http.StatusOK, "gin OK")
	})

	// Get counter
	app.GET("/count", func(c *gin.Context) {
		c.String(http.StatusOK, "%d", counter)
	})

	// Increment counter
	app.POST("/inc", func(c *gin.Context) {
		counter++
		c.String(http.StatusOK, "%d", counter)
	})

	// Decrement counter
	app.POST("/dec", func(c *gin.Context) {
		counter--
		c.String(http.StatusOK, "%d", counter)
	})

	return app
}

func main() {
	r := setupRouter()
	r.Run(":5000")
	app := setupHTTPServer()
	app.Run(":5000")
}
