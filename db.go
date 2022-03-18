package main

import (
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
