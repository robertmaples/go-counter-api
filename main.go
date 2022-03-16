package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type counter struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

var counters = []counter{
	{ID: "5ca44aab-ee12-4911-925c-329175c0d1a0", Value: 50},
	{ID: "d09b11a1-3ef8-47f6-a4de-620e7cabdc1a", Value: 100},
	{ID: "d09b11a1-3ef8-47f6-a4de-620e7cabdc1a", Value: 200},
}

// getCounters responds with the list of all counters as JSON.
func getCounters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, counters)
}

func main() {
	router := gin.Default()
	router.GET("/counters", getCounters)

	router.Run("localhost:8080")
}
