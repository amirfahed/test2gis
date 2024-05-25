package handlers

import (
	"encoding/json"
	"net/http"
	"applicationDesignTest/internal/services"
	"applicationDesignTest/internal/models"
	"applicationDesignTest/internal/utils/logger"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	success, unavailableDays := services.CreateOrder(newOrder)
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
