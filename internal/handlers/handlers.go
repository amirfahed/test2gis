package handlers

import (
	"applicationDesignTest/internal/models"
	"applicationDesignTest/internal/service"
	"applicationDesignTest/internal/utils/helpers"
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
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	errMsg, ok := helpers.ValidateCreateOrderRequest(newOrder)
	if !ok {
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	success, unavailableDays := handler.orderService.CreateOrder(ctx, newOrder)
	if !success {
		http.Error(w, "Hotel room is not available for selected dates", http.StatusConflict)
		logger.LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newOrder)

	logger.LogInfo("Order successfully created: %v", newOrder)
}
