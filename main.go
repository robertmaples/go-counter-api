package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// TODO: make id unique value
type counter struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

var counters = []counter{
	{ID: "5ca44aab-ee12-4911-925c-329175c0d1a0", Value: 50},
	{ID: "d09b11a1-3ef8-47f6-a4de-620e7cabdc1a", Value: 100},
	{ID: "d09b11a1-3ef8-47f6-a4de-620e7cabdc1a", Value: 200},
}

func main() {
	router := gin.Default()

	router.GET("/counters", getCounters)
	router.GET("/counter/:id", getCounterByID)
	router.POST("/counter", postCounter)
	router.POST("/counter/:id", incrementCounter)
	router.DELETE("/counter/:id", deleteCounter)

	router.Run("localhost:8080")
}

// getCounters responds with the list of all counters as JSON.
func getCounters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, counters)
}

// postCounter adds a counter from JSON received in the request body.
func postCounter(c *gin.Context) {
	id := uuid.New()

	newCounter := counter{ID: id.String(), Value: 0}

	counters = append(counters, newCounter)
	c.IndentedJSON(http.StatusCreated, newCounter)
}

// getCounterByID locates the counter with the id sent in by the request.
func getCounterByID(c *gin.Context) {
	id := c.Param("id")

	for _, counter := range counters {
		if counter.ID == id {
			c.IndentedJSON(http.StatusOK, counter)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}

// incrementCounter locates the counter with the id sent in by the request and increments
// its value by one.
func incrementCounter(c *gin.Context) {
	id := c.Param("id")

	for _, counter := range counters {
		if counter.ID == id {
			counter.Value += 1
			c.IndentedJSON(http.StatusOK, counter)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}

// deleteCounter locates the counter with the id sent in by the request and deletes it.
func deleteCounter(c *gin.Context) {
	id := c.Param("id")

	for i, counter := range counters {
		if counter.ID == id {
			counters = append(counters[:i], counters[i+1:]...)
			c.IndentedJSON(http.StatusOK, counter)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}
