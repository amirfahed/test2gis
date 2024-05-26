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
	"applicationDesignTest/internal/utils/logger"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	logger.LogInfo("Server listening on localhost:8080")
	conf, _ := config.New()
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Host, conf.Port), router.GetRouter())
	if errors.Is(err, http.ErrServerClosed) {
		logger.LogInfo("Server closed")
	} else if err != nil {
		logger.LogErrorf("Server error: %s", err)
		os.Exit(1)
	}
}
