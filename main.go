package main

import (
	"log"
)

func main() {
	db, err := dbInit()

	if err != nil {
		log.Println(err.Error())
	}

	s := &server{
		db: db,
	}

	router := setupRouter(s)

	router.LoadHTMLGlob("templates/*")

	routerErr := router.Run()
	if routerErr != nil {
		return
	}
}
