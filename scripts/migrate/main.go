package main

import database "main/database"

func main() {
	db := database.GetDB()

	err := db.Migrator().DropTable(&database.User{}, &database.ShareGroup{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&database.User{}, &database.ShareGroup{})
	if err != nil {
		panic(err)
	}
}
