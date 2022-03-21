package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TODO: Make id unique
type Counter struct {
	gorm.Model
	ID    string `gorm:"uniqueIndex" json:"id"`
	Value int    `json:"value"`
}

func dbInit() *gorm.DB {
	db, dbErr := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if dbErr != nil {
		panic("failed to connect database")
	}

	return db
}

func dbGetCounters() []Counter {
	db := dbInit()

	var counters []Counter
	db.Find(&counters)

	return counters
}

func dbGetCounterById(id string) Counter {
	db := dbInit()

	var counter Counter
	db.Where("id = ?", id).Find(&counter)

	return counter
}

func dbCreateCounter() Counter {
	db := dbInit()

	id := uuid.New()
	newCounter := Counter{ID: id.String(), Value: 0}

	db.Create(&newCounter)

	return newCounter
}

func dbIncrementCounter(id string) Counter {
	db := dbInit()

	var counter Counter
	db.Where("id = ?", id).Find(&counter).Update("value", counter.Value+1)

	return counter
}

func dbDeleteCounter(id string) Counter {
	db := dbInit()

	var counter Counter
	db.Where("id = ?", id).Find(&counter).Delete(&counter)

	return counter
}
