package service

import (
	"applicationDesignTest/internal/models"
	"applicationDesignTest/internal/repository"
	"applicationDesignTest/internal/utils/helpers"
	"applicationDesignTest/internal/utils/logger"
	"context"
	"time"
)

type OrderService struct {
	orderRepo repository.OrderRepo
	hotelRepo repository.HotelRepo
}

func NewOrderService(orderRepo repository.OrderRepo, hotelRepo repository.HotelRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		hotelRepo: hotelRepo,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, newOrder models.Order) (bool, map[time.Time]struct{}) {
	daysToBook := helpers.DaysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	availability, err := o.hotelRepo.GetRoomAvailability(ctx)
	if err != nil {
		logger.LogErrorf(err.Error())
	}

	var availabilityToUpdate []models.RoomAvailability
	for _, dayToBook := range daysToBook {
		for _, availability := range availability {
			if !availability.Date.Equal(dayToBook) ||
				availability.Quota < 1 ||
				availability.HotelID != newOrder.HotelID ||
				availability.RoomID != newOrder.RoomID {
				continue
			}
			availability.Quota -= 1
			availabilityToUpdate = append(availabilityToUpdate, availability)
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		return false, unavailableDays
	}

	err = o.hotelRepo.UpdateRoomAvailability(ctx, availabilityToUpdate)
	if err != nil {
		logger.LogErrorf(err.Error())
		return false, nil
	}

	err = o.orderRepo.CreateOrders(ctx, []models.Order{newOrder})
	if err != nil {
		return false, nil
	}

	if err != nil {
		logger.LogErrorf(err.Error())
		return false, nil
	}

	return true, nil
}
