package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/counters", getCounters)
	r.GET("/counter/:id", getCounterByID)
	r.POST("/counter", createCounter)
	r.POST("/counter/:id", incrementCounter)
	r.DELETE("/counter/:id", deleteCounter)

	return r
}
