package crud

import (
	"errors"
	"log"
	"main/database"
	"time"

	"gorm.io/gorm"
)

func GetUserById(db *gorm.DB, userId int) *database.User {
	var user database.User
	res := db.Preload("ShareGroup").First(&user, userId)

	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			log.Println(res.Error)
		}

		return nil
	}

	return &user
}

func CreateUser(db *gorm.DB, name string, iconID string) (
	*database.User, error,
) {
	user := database.User{Name: name, IconID: iconID}

	if res := db.Create(&user); res.Error != nil {
		log.Printf("Error: %v\n", res.Error)
		return nil, res.Error
	}

	return &user, nil
}

func UpdatePosition(db *gorm.DB, user *database.User, lat float64, lon float64) (
	*database.User, error,
) {
	user.Lat = &lat
	user.Lon = &lon
	now := time.Now()
	user.PositionAt = &now

	if res := db.Save(&user); res.Error != nil {
		log.Printf("Error: %v\n", res.Error)
		return nil, res.Error
	}

	return user, nil
}

func UpdateIsArrived(db *gorm.DB, user *database.User) (*database.User, error) {
	user.IsArrived = true

	if res := db.Save(&user); res.Error != nil {
		log.Printf("Error: %v\n", res.Error)
		return nil, res.Error
	}

	return user, nil
}

func UpdateShortMessage(db *gorm.DB, user *database.User, shortMessage *string) (
	*database.User, error,
) {
	user.ShortMessage = shortMessage

	if res := db.Save(&user); res.Error != nil {
		log.Printf("Error: %v\n", res.Error)
		return nil, res.Error
	}

	return user, nil
}
