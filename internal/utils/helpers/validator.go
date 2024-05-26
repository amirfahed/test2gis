package helpers

import (
	"applicationDesignTest/internal/models"
	"net/mail"
)

func ValidateCreateOrderRequest(orderRequest models.Order) (string, bool) {
	if !isValidEmail(orderRequest.UserEmail) {
		return "Email is not valid", false
	}
	if orderRequest.RoomID == "" {
		return "RoomID is required", false
	}

	if orderRequest.HotelID == "" {
		return "HotelID is required", false
	}
	if orderRequest.From.After(orderRequest.To) {
		return "From date after To date", false
	}
	return "", true
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
