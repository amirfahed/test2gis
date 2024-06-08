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
	"applicationDesignTest/internal/models"
	"applicationDesignTest/internal/router"
	"applicationDesignTest/internal/utils/logger"
	"errors"
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger.LogInfo("Start server")
	conf, _ := config.New()
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.AppConfig.DB.Host,
		config.AppConfig.DB.User,
		config.AppConfig.DB.Password,
		config.AppConfig.DB.DBName,
		config.AppConfig.DB.Port,
	)

	// Подключение к базе данных PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogErrorf("failed to connect database")
	}

	// Автомиграция таблицы usersdb
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		logger.LogErrorf("failed to migrate database")
	}

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port), router.GetRouter())
	if errors.Is(err, http.ErrServerClosed) {
		logger.LogInfo("Server closed")
	} else if err != nil {
		logger.LogErrorf("Server error: %s", err)
		os.Exit(1)
	}

}
