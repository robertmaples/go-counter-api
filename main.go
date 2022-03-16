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
	{ID: "fbe31350-31db-4117-9a60-4e33eb184f65", Value: 200},
}

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

func main() {
	router := setupRouter()

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

// getCounters responds with the list of all counters as JSON.
func getCounters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, counters)
}

// createCounter adds a counter from JSON received in the request body.
func createCounter(c *gin.Context) {
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

	for i, count := range counters {
		if count.ID == id {
			counters[i].Value++
			c.IndentedJSON(http.StatusOK, counters[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}

// deleteCounter locates the counter with the id sent in by the request and deletes it.
func deleteCounter(c *gin.Context) {
	id := c.Param("id")

	for i, count := range counters {
		if count.ID == id {
			counters = append(counters[:i], counters[i+1:]...)
			c.IndentedJSON(http.StatusOK, count)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}
