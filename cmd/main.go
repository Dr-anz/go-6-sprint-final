package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер с базовыми настройками
	logger := log.New(
		os.Stdout,     // куда писать логи (в консоль)
		"SERVER: ",    // префикс для всех сообщений
		log.LstdFlags, // стандартные флаги (дата, время)
	)

	// Создаем сервер, передавая созданный логгер
	srv := server.NewServer(logger)

	// Запускаем сервер
	err := srv.Start()
	if err != nil {
		// Если произошла ошибка при запуске - выводим фатальную ошибку
		logger.Fatal(err)
	}
}
