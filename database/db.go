package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	ShortMessage *string
	ShareGroupID *uint64
	ShareGroup   *ShareGroup
	Lat          *float64
	Lon          *float64
	PositionAt   *time.Time
	IsArrived    bool `gorm:"default:false"`
}

type ShareGroup struct {
	gorm.Model
	LinkKey     string
	DestLon     float64
	DestLat     float64
	MeetingTime string
	Address     string
	AdminUserID uint64
	AdminUser   User
	Users       []*User
}

var globalDB *gorm.DB

func GetDB() *gorm.DB {
	if globalDB != nil {
		return globalDB
	}

	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	globalDB = db

	return db
}
