package repository

import (
	"applicationDesignTest/internal/models"
	"context"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context) error
}

type HotelRepo interface {
	GetRooms(ctx context.Context) ([]models.RoomAvailability, error)
	ReserveRooms(ctx context.Context, rooms []models.RoomAvailability) error
}
