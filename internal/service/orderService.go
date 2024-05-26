package service

import (
	"applicationDesignTest/internal/models"
	"applicationDesignTest/internal/utils/helpers"
	"context"
	"time"
)

var Orders = []models.Order{}
var Availability = []models.RoomAvailability{
	{"reddison", "lux", helpers.Date(2024, 1, 1), 1},
	{"reddison", "lux", helpers.Date(2024, 1, 2), 1},
	{"reddison", "lux", helpers.Date(2024, 1, 3), 1},
	{"reddison", "lux", helpers.Date(2024, 1, 4), 1},
	{"reddison", "lux", helpers.Date(2024, 1, 5), 0},
}

type OrderService struct {
	//repoOrder := orderRepo.NewOrderRepo()
	//repoRoom := roomRepo.NewRoomRepo()
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) CreateOrder(ctx context.Context, newOrder models.Order) (bool, map[time.Time]struct{}) {
	daysToBook := helpers.DaysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	for _, dayToBook := range daysToBook {
		for i, availability := range Availability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			Availability[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		return false, unavailableDays
	}

	Orders = append(Orders, newOrder)
	return true, nil
}
