package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"robert": "password",
	}))

	authorized.GET("/counters", getCounters)
	authorized.GET("/counter/:id", getCounterByID)
	authorized.POST("/counter", createCounter)
	authorized.POST("/counter/:id", incrementCounter)
	authorized.DELETE("/counter/:id", deleteCounter)

	return r
}
