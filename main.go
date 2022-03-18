package main

func main() {
	db := dbInit()

	// Migrate the schema
	db.AutoMigrate(&[]Counter{})

	router := setupRouter()
	routerErr := router.Run("localhost:8080")
	if routerErr != nil {
		return
	}
}
