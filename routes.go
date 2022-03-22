package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// showIndexPage returns all counters
func (s server) index(c *gin.Context) {
	counters, err := s.dbGetCounters()

	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "index.html", gin.H{"title": "Home Page", "payload": counters})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"payload": counters})
}

// getCounters responds with the list of all counters as JSON.
func (s server) getCounters(c *gin.Context) {
	counters, err := s.dbGetCounters()

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not fulfill request."})
	}

	c.IndentedJSON(http.StatusOK, counters)
}

// createCounter adds a counter from JSON received in the request body.
func (s server) createCounter(c *gin.Context) {
	counter, err := s.dbCreateCounter()

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not fulfill request."})
	}

	c.IndentedJSON(http.StatusCreated, counter)
}

// getCounterByID locates the counter with the id sent in by the request.
func (s server) getCounterByID(c *gin.Context) {
	id := c.Param("id")

	counter, err := s.dbGetCounterById(id)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
	}

	c.IndentedJSON(http.StatusOK, counter)
}

// incrementCounter locates the counter with the id sent in by the request and increments
// its value by one.
func (s server) incrementCounter(c *gin.Context) {
	id := c.Param("id")

	counter, err := s.dbIncrementCounter(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": "counter not found"})
	}

	c.IndentedJSON(http.StatusOK, counter)
}

// deleteCounter locates the counter with the id sent in by the request and deletes it.
func (s server) deleteCounter(c *gin.Context) {
	id := c.Param("id")

	counter, err := s.dbDeleteCounter(id)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
	}

	c.IndentedJSON(http.StatusOK, counter)
}
