package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type server struct {
	db *gorm.DB
}

type Counter struct {
	gorm.Model
	ID    string `gorm:"uniqueIndex" json:"id"`
	Value int    `json:"value"`
}

func dbInit() (db *gorm.DB, err error) {
	db, dbErr := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// Migrate the schema
	if migrationErr := db.AutoMigrate(&[]Counter{}); migrationErr != nil {
		log.Printf("DB Migration Error: %v", migrationErr)
		return nil, migrationErr
	}

	return db, dbErr
}

func (s server) dbGetCounters() ([]Counter, error) {
	var counters []Counter

	if res := s.db.Find(&counters); res.Error != nil {
		return nil, res.Error
	}

	return counters, nil
}

func (s server) dbGetCounterById(id string) (Counter, error) {
	var counter Counter

	if res := s.db.Where("id = ?", id).Find(&counter); res.Error != nil {
		return counter, res.Error
	}

	return counter, nil
}

func (s server) dbCreateCounter() (Counter, error) {
	id := uuid.New()

	newCounter := Counter{ID: id.String(), Value: 0}

	if res := s.db.Create(&newCounter); res.Error != nil {
		return newCounter, res.Error
	}

	return newCounter, nil
}

func (s server) dbIncrementCounter(id string) (Counter, error) {
	var counter Counter

	if res := s.db.Where("id = ?", id).Find(&counter).Update("value", counter.Value+1); res.Error != nil {
		return counter, res.Error
	}

	return counter, nil
}

func (s server) dbDeleteCounter(id string) (Counter, error) {
	var counter Counter

	if res := s.db.Where("id = ?", id).Find(&counter).Delete(&counter); res.Error != nil {
		return counter, res.Error
	}

	return counter, nil
}
