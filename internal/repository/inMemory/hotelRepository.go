package inMemory

import (
	"applicationDesignTest/internal/models"
	"applicationDesignTest/internal/utils/helpers"
	"context"
)

type HotelRepository struct {
	roomAvailabilities []models.RoomAvailability
}

func NewHotelRepository() *HotelRepository {
	initialRooms := []models.RoomAvailability{
		{"reddison", "lux", helpers.Date(2024, 1, 1), 1},
		{"reddison", "lux", helpers.Date(2024, 1, 2), 1},
		{"reddison", "lux", helpers.Date(2024, 1, 3), 1},
		{"reddison", "lux", helpers.Date(2024, 1, 4), 1},
		{"reddison", "lux", helpers.Date(2024, 1, 5), 0},
	}
	return &HotelRepository{
		roomAvailabilities: initialRooms,
	}
}

func (hotel *HotelRepository) GetRoomAvailability(ctx context.Context) ([]models.RoomAvailability, error) {
	return hotel.roomAvailabilities, nil
}

func (hotel *HotelRepository) UpdateRoomAvailability(ctx context.Context, updatedRooms []models.RoomAvailability) error {
	for _, updatedRoom := range updatedRooms {
		for _, roomAvailability := range hotel.roomAvailabilities {
			if roomAvailability.IsEqualRoom(updatedRoom) {
				roomAvailability = updatedRoom
			}
		}
	}

	return nil
}
