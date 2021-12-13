package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
)

var db = make(map[string]string)
var redisContext = context.Background()

func setupHTTPServer(rdb *redis.Client) *gin.Engine {
	app := gin.Default()

	err := rdb.Set(redisContext, "count", 0, 0).Err()
	if err != nil {
		panic(err)
	}

	// Get OK
	app.GET("/ok", func(c *gin.Context) {
		c.String(http.StatusOK, "gin OK")
	})

	// Get counter
	app.GET("/count", func(c *gin.Context) {
		sVal, err := rdb.Get(redisContext, "count").Result()
		if err != nil {
			panic(err)
		}
		iVal, err := strconv.Atoi(sVal)
		if err != nil {
			panic(err)
		}
		c.String(http.StatusOK, "%d", iVal)
	})

	// Increment counter
	app.POST("/inc", func(c *gin.Context) {
		val, err := rdb.Incr(redisContext, "count").Result()
		if err != nil {
			panic(err)
		}
		c.String(http.StatusOK, "%d", val)
	})

	// Decrement counter
	app.POST("/dec", func(c *gin.Context) {
		val, err := rdb.Decr(redisContext, "count").Result()
		if err != nil {
			panic(err)
		}
		c.String(http.StatusOK, "%d", val)
	})

	return app
}

func setupRedisClient(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(redisContext, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(redisContext, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(redisContext, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	return rdb
}

func main() {
	rdb := setupRedisClient("redis:6379")

	app := setupHTTPServer(rdb)
	app.Run(":5000")
}
