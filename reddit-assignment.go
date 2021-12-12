package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var db = make(map[string]string)
var redisContext = context.Background()

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

func setupRedisClient(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Set(redisContext, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(redisContext, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get(redisContext, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	return client
}

func main() {
	_ = setupRedisClient("localhost:6379")

	app := setupHTTPServer()
	app.Run(":5000")
}
