package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Counter struct {
	ID    string `gorm:"uniqueIndex" json:"id"`
	Value int    `json:"value"`
}

func dbInit() error {
	var dbErr error
	db, dbErr = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&[]Counter{})

	return dbErr
}

func dbGetCounters() ([]Counter, error) {
	var counters []Counter
	res := db.Find(&counters)
	if res.Error != nil {
		return counters, res.Error
	}

	return counters, nil
}

func dbGetCounterById(id string) (Counter, error) {
	var counter Counter
	if res := db.Where("id = ?", id).Find(&counter); res.Error != nil {
		return counter, res.Error
	}

	return counter, nil
}

func dbCreateCounter() (Counter, error) {
	id := uuid.New()
	newCounter := Counter{ID: id.String(), Value: 0}

	if res := db.Create(&newCounter); res.Error != nil {
		return newCounter, res.Error
	}

	return newCounter, nil
}

func dbIncrementCounter(id string) (Counter, error) {
	var counter Counter
	if res := db.Where("id = ?", id).Find(&counter).Update("value", counter.Value+1); res.Error != nil {
		return counter, res.Error
	}

	return counter, nil
}

func dbDeleteCounter(id string) (Counter, error) {
	var counter Counter
	res := db.Where("id = ?", id).Find(&counter).Delete(&counter)
	if res.Error != nil {
		return counter, res.Error
	}

	return counter, nil
}
