package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server структура для хранения конфигурации сервера
type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

// NewServer создает и настраивает HTTP сервер
func NewServer(logger *log.Logger) *Server {
	// Создаем роутер
	router := http.NewServeMux()

	// Регистрируем хендлеры
	router.HandleFunc("/", handlers.RootHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	// Создаем и настраиваем сервер
	server := &Server{
		logger: logger,
		httpServer: &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

	return server
}

// Start запускает сервер
func (s *Server) Start() error {
	s.logger.Printf("Запуск сервера на порту 8080...")
	return s.httpServer.ListenAndServe()
}
