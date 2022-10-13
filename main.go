package main

import (
	"assignment-2/database"
	"assignment-2/routers"
	"fmt"
)

func main() {
	database.StartDB()
	routers.StartServer().Run()

	fmt.Println("Successfully connected to database")
}
