package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// showIndexPage returns all counters
func showIndexPage(c *gin.Context) {
	counters, err := dbGetCounters()

	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "index.html", gin.H{"title": "Home Page", "payload": counters})
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"payload": counters})
}

// getCounters responds with the list of all counters as JSON.
func getCounters(c *gin.Context) {
	counters, err := dbGetCounters()

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not fulfill request."})
	}

	c.IndentedJSON(http.StatusOK, counters)
}

// createCounter adds a counter from JSON received in the request body.
func createCounter(c *gin.Context) {
	counter, err := dbCreateCounter()

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not fulfill request."})
	}

	c.IndentedJSON(http.StatusCreated, counter)
}

// getCounterByID locates the counter with the id sent in by the request.
func getCounterByID(c *gin.Context) {
	id := c.Param("id")

	counter, err := dbGetCounterById(id)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
	}

	c.IndentedJSON(http.StatusOK, counter)
}

// incrementCounter locates the counter with the id sent in by the request and increments
// its value by one.
func incrementCounter(c *gin.Context) {
	id := c.Param("id")

	counter, err := dbIncrementCounter(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": "counter not found"})
	}

	c.IndentedJSON(http.StatusOK, counter)
}

// deleteCounter locates the counter with the id sent in by the request and deletes it.
func deleteCounter(c *gin.Context) {
	id := c.Param("id")

	counter, err := dbDeleteCounter(id)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
	}

	c.IndentedJSON(http.StatusOK, counter)
}
