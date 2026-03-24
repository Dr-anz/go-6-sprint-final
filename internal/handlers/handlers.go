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
	// 1. Парсим multipart-форму (максимум 10 МБ)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Ошибка парсинга multipart-формы: %v", err)
		http.Error(w, "Ошибка парсинга формы", http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Файл не найден в форме: %v. Доступные поля: %v",
			err, r.MultipartForm.Value)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка получения файла: поле 'file' отсутствует"))
		return
	}
	defer file.Close()

	log.Printf("Получен файл: %s, размер: %d байт",
		header.Filename, header.Size)

	// Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	inputText := string(data)
	if inputText == "" {
		log.Printf("Получен пустой файл")
		http.Error(w, "Файл пуст", http.StatusBadRequest)
		return
	}

	// Передаём данные в функцию автоопределения из пакета service
	convertedText, err := service.Convert(inputText)
	if err != nil {
		log.Printf("Ошибка конвертации: %v", err)
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}

	// Генерируем имя файла с использованием времени и расширения исходного файла
	timestamp := time.Now().UTC().String()
	fileExt := filepath.Ext(header.Filename)
	outputFilename := "converted_" + timestamp + fileExt

	// Создаём и записываем в локальный файл
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		log.Printf("Ошибка создания файла %s: %v", outputFilename, err)
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(convertedText)
	if err != nil {
		log.Printf("Ошибка записи в файл %s: %v", outputFilename, err)
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат конвертации строки
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedText))
}
