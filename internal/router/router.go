package router

import (
	"applicationDesignTest/internal/handlers"
	"applicationDesignTest/internal/repository/inMemory"
	"applicationDesignTest/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	handler := handlers.NewOrderHandler(service.NewOrderService(inMemory.NewOrderRepository(), inMemory.NewHotelRepository()))
	r.Route("/orders", func(r chi.Router) {
		r.Post("/", handler.CreateOrder)
	})

	return r
}
