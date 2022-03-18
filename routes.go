package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

// getCounters responds with the list of all counters as JSON.
func getCounters(c *gin.Context) {
	db := dbInit()

	var counters []Counter
	db.Find(&counters)
	fmt.Println("{}", counters)

	c.IndentedJSON(http.StatusOK, counters)
}

// createCounter adds a counter from JSON received in the request body.
func createCounter(c *gin.Context) {
	db := dbInit()

	id := uuid.New()
	newCounter := Counter{ID: id.String(), Value: 0}

	db.Create(&newCounter)

	c.IndentedJSON(http.StatusCreated, newCounter)
}

// getCounterByID locates the counter with the id sent in by the request.
func getCounterByID(c *gin.Context) {
	db := dbInit()

	id := c.Param("id")

	var counter Counter
	db.Where("id = ?", id).Find(&counter)

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
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("failed to connect to database")
	}

	id := c.Param("id")

	var counter Counter
	db.Where("id = ?", id).Find(&counter).Update("value", counter.Value+1)

	if counter.ID == id {
		c.IndentedJSON(http.StatusOK, counter)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}

// deleteCounter locates the counter with the id sent in by the request and deletes it.
func deleteCounter(c *gin.Context) {
	db := dbInit()

	id := c.Param("id")

	var counter Counter
	db.Where("id = ?", id).Find(&counter).Delete(&counter)
	if counter.ID == id {
		c.IndentedJSON(http.StatusOK, counter)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "counter not found"})
}
