package repository

import (
	"applicationDesignTest/internal/models"
	"context"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, orders []models.Order) error
}

type HotelRepo interface {
	GetRoomAvailability(ctx context.Context) ([]models.RoomAvailability, error)
	UpdateRoomAvailability(ctx context.Context, updatedRoom []models.RoomAvailability) error
}
