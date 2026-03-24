package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет тип ввода (текст или код Морзе)
// и выполняет соответствующую конвертацию с использованием пакета morse
// Если передан текст — конвертирует в код Морзе через morse.ToMorse
// Если передан код Морзе — конвертирует в текст через morse.ToText
// Возвращает результат и ошибку (если есть)
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("входная строка пуста")
	}

	// Определяем, является ли ввод кодом Морзе:
	// если строка содержит только '.', '-' — считаем, что это Морзе
	isMorse := true
	for _, char := range input {
		if char != '.' && char != '-' {
			isMorse = false
			break
		}
	}

	if isMorse {
		result := morse.ToText(input)
		// Проверяем, не получился ли пустой результат — это может указывать на ошибку в коде Морзе
		if result == "" {
			return "", errors.New("неверный код Морзе")
		}
		return result, nil
	} else {
		result := morse.ToMorse(input)
		// Проверяем, не получился ли пустой результат — это может указывать на неподдерживаемые символы
		if result == "" {
			return "", errors.New("текст содержит неподдерживаемые символы")
		}
		return result, nil
	}
}
