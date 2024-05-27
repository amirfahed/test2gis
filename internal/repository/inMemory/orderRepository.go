package inMemory

import (
	"applicationDesignTest/internal/models"
	"context"
	"sync"
)

type OrderRepository struct {
	orders []models.Order
	mu     sync.RWMutex
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make([]models.Order, 0),
	}
}

func (orderRepo *OrderRepository) CreateOrders(ctx context.Context, orders []models.Order) error {
	orderRepo.mu.Lock()
	defer orderRepo.mu.Unlock()
	orderRepo.orders = append(orderRepo.orders, orders...)
	return nil
}
