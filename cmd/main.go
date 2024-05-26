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
	"applicationDesignTest/config"
	"applicationDesignTest/internal/router"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Start server")
	conf, _ := config.New()
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Host, conf.Port), router.GetRouter())
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed")
	} else if err != nil {
		log.Fatalf("Server error: %s", err)
		os.Exit(1)
	}
}
