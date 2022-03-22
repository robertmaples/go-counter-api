package main

import "github.com/gin-gonic/gin"

func setupRouter(s *server) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"robert": "password",
	}))

	authorized.GET("/", s.index)
	authorized.GET("/counters", s.getCounters)
	authorized.GET("/counter/:id", s.getCounterByID)
	authorized.POST("/counter", s.createCounter)
	authorized.POST("/counter/:id", s.incrementCounter)
	authorized.DELETE("/counter/:id", s.deleteCounter)

	return r
}
