// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"applicationDesignTest/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/orders", func(r chi.Router) {
		r.Post("/", handlers.CreateOrder)
	})

	log.Println("Сервер слушает на localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Сервер закрыт")
	} else if err != nil {
		log.Fatalf("Ошибка сервера: %s", err)
		os.Exit(1)
	}
}
