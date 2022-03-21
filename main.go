package main

func main() {
	db := dbInit()

	// Migrate the schema
	db.AutoMigrate(&[]Counter{})

	router := setupRouter()

	router.LoadHTMLGlob("templates/*")

	routerErr := router.Run()
	if routerErr != nil {
		return
	}
}
