package crud

import (
	"log"
	"main/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func genUUIDV4() string {
	uuidv4, _ := uuid.NewRandom()
	return uuidv4.String()
}

func CreateShareGroup(db *gorm.DB, destLng float64, destLat float64, meetingTime string) (
	*database.ShareGroup, error,
) {
	shareGroup := database.ShareGroup{
		LinkKey:     genUUIDV4(),
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
