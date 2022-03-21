package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
*	TODO: Create handlers to separate model/db from request response
 */
// show index page of conters
func showIndexPage(c *gin.Context) {
	counters := dbGetCounters()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"payload": counters})
}

// getCounters responds with the list of all counters as JSON.
func getCounters(c *gin.Context) {
	counters := dbGetCounters()

	c.IndentedJSON(http.StatusOK, counters)
}

// createCounter adds a counter from JSON received in the request body.
func createCounter(c *gin.Context) {
	newCounter := dbCreateCounter()

	c.IndentedJSON(http.StatusCreated, newCounter)
}

// getCounterByID locates the counter with the id sent in by the request.
func getCounterByID(c *gin.Context) {
	id := c.Param("id")

	counter := dbGetCounterById(id)

	// TODO: How best to confirm correct counter is found and return proper status
	if counter.ID == id {
		c.IndentedJSON(http.StatusOK, counter)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}

// incrementCounter locates the counter with the id sent in by the request and increments
// its value by one.
func incrementCounter(c *gin.Context) {
	id := c.Param("id")

	counter := dbIncrementCounter(id)

	if counter.ID == id {
		c.IndentedJSON(http.StatusOK, counter)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}

// deleteCounter locates the counter with the id sent in by the request and deletes it.
func deleteCounter(c *gin.Context) {

	id := c.Param("id")
	counter := dbDeleteCounter(id)

	if counter.ID == id {
		c.IndentedJSON(http.StatusOK, counter)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}
