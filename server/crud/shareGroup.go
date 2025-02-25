package crud

import (
	"errors"
	"log"
	"main/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenUUIDV4() string {
	uuidv4, _ := uuid.NewRandom()
	return uuidv4.String()
}

func GetShareGroupByLinkKey(db *gorm.DB, linkKey string) *database.ShareGroup {
	var shareGroup database.ShareGroup
	err := db.Where("link_key = ?", linkKey).Preload("Users").First(&shareGroup).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("error: %v\n", err)
		}

		return nil
	}

	return &shareGroup
}

func CreateShareGroup(db *gorm.DB, destLon float64, destLat float64, meetingTime string) (
	*database.ShareGroup, error,
) {
	shareGroup := database.ShareGroup{
		LinkKey:     GenUUIDV4(),
		DestLon:     destLon,
		DestLat:     destLat,
		MeetingTime: meetingTime,
	}
	if res := db.Create(&shareGroup); res.Error != nil {
		log.Printf("Error: %v\n", res.Error)
		return nil, res.Error
	}

	return &shareGroup, nil
}

func JoinShareGroup(db *gorm.DB, shareGroup *database.ShareGroup, user *database.User) (
	*database.ShareGroup, error,
) {
	user.ShareGroup = shareGroup

	err := db.Save(user).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}

	updatedShareGroup := GetShareGroupByLinkKey(db, shareGroup.LinkKey)

	return updatedShareGroup, nil
}
