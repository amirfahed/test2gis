package handlers

import (
	"applicationDesignTest/internal/models"
	"applicationDesignTest/internal/service"
	"applicationDesignTest/internal/utils/logger"
	"context"
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	orderService service.Service
}

func NewOrderHandler(orderService service.Service) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	var newOrder models.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	success, unavailableDays := handler.orderService.CreateOrder(ctx, newOrder)
	if !success {
		http.Error(w, "Номера недоступны на выбранные даты", http.StatusConflict)
		logger.LogErrorf("Номера недоступны на выбранные даты:\n%v\n%v", newOrder, unavailableDays)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)

	logger.LogInfo("Заказ успешно создан: %v", newOrder)
}
