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
	ShareGroupID *uint64
	ShareGroup   *ShareGroup
	Lat          *float64
	Lon          *float64
	PositionAt   *time.Time
}

type ShareGroup struct {
	gorm.Model
	LinkKey     string
	DestLon     float64
	DestLat     float64
	MeetingTime string
	Users       []*User
}

func GetDB() *gorm.DB {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}
	return db
}
