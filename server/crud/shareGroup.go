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
	res := db.Where("linkkey = ?", linkKey).First(&shareGroup)

	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			log.Println(res.Error)
		}

		return nil
	}

	return &shareGroup
}

func CreateShareGroup(db *gorm.DB, destLng float64, destLat float64, meetingTime string) (
	*database.ShareGroup, error,
) {
	shareGroup := database.ShareGroup{
		LinkKey:     GenUUIDV4(),
		DestLng:     destLng,
		DestLat:     destLat,
		MeetingTime: meetingTime,
	}
	if res := db.Create(&shareGroup); res.Error != nil {
		log.Fatalf("Error: %v", res.Error)
		return nil, res.Error
	}

	return &shareGroup, nil
}

func JoinShareGroupByLinkKey(db *gorm.DB, user database.User, shareGroup database.ShareGroup) (
	*database.ShareGroup, error,
) {
	user.ShareGroupID = int(shareGroup.ID)

	if res := db.Save(&user); res.Error != nil {
		log.Fatalf("Error: %v", res.Error)
		return nil, res.Error
	}

	return &shareGroup, nil
}
