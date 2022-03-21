package main

import (
	"log"
)

func main() {
	err := dbInit()

	if err != nil {
		log.Println(err.Error())
	}

	router := setupRouter()

	router.LoadHTMLGlob("templates/*")

	routerErr := router.Run()
	if routerErr != nil {
		return
	}
}
