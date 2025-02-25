package database

import (
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
	DestLng     float64
	DestLat     float64
	MeetingTime string
	Users       []*User
}

func GetDB() *gorm.DB {
	dsn := "host=localhost user=user password=password dbname=db port=15000 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}
	return db
}
