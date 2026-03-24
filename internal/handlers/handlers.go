package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// RootHandler возвращает HTML из файла index.html для корневого эндпоинта
func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файла и конвертацию данных
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем максимальный размер загружаемого файла (10 МБ)
	r.ParseMultipartForm(10 << 20)

	// Получаем файл из формы
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		log.Printf("Ошибка получения файла: %v", err)
		return
	}
	defer file.Close() // Обязательно закрываем файл после использования

	// Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		log.Printf("Ошибка чтения файла: %v", err)
		return
	}

	inputText := string(data)

	// Передаём данные в функцию автоопределения из пакета service
	convertedText, err := service.Convert(inputText)
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Ошибка конвертации: %v", err)
		return
	}

	// Генерируем имя файла с использованием времени
	timestamp := time.Now().UTC().String()
	fileExt := filepath.Ext(header.Filename)
	outputFilename := "converted_" + timestamp + fileExt

	// Создаём и записываем в локальный файл
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		log.Printf("Ошибка создания файла: %v", err)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(convertedText)
	if err != nil {
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		log.Printf("Ошибка записи в файл: %v", err)
		return
	}

	// Возвращаем результат конвертации строки
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(convertedText))
}
