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
