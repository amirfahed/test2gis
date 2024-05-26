package service

import (
	"applicationDesignTest/internal/models"
	"context"
	"time"
)

type Service interface {
	CreateOrder(ctx context.Context, newOrder models.Order) (bool, map[time.Time]struct{})
}
