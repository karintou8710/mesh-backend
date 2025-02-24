package main

import database "main/database"

func main() {
	db := database.GetDB()

	db.AutoMigrate(&database.User{})
}
