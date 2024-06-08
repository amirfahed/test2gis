package models

import "time"
import "gorm.io/gorm"


type Order struct {
	gorm.Model
	HotelID   string    `gorm:"not null" json:"hotel_id"`
	RoomID    string    `gorm:"not null" json:"room_id"`
	UserEmail string    `gorm:"not null" json:"email"`
	From      time.Time `gorm:"not null json:"from"`
	To        time.Time `gorm:"not null json:"to"`
}
