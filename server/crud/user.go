package crud

import (
	"errors"
	"log"
	"main/database"

	"gorm.io/gorm"
)

func GetUserById(db *gorm.DB, userId int) *database.User {
	var user database.User
	res := db.First(&user, userId)

	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			log.Println(res.Error)
		}

		return nil
	}

	return &user
}

func CreateUser(db *gorm.DB, name string) (
	*database.User, error,
) {
	user := database.User{Name: name}

	if res := db.Create(&user); res.Error != nil {
		log.Fatalf("Error: %v", res.Error)
		return nil, res.Error
	}

	return &user, nil
}
