package models

import "time"

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}

func (hotel *RoomAvailability) IsEqualRoom(roomAvailability RoomAvailability) bool {
	return hotel.HotelID == roomAvailability.HotelID && hotel.RoomID == roomAvailability.RoomID && hotel.Date == roomAvailability.Date
}
