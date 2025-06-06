package randstring

import (
	"math/rand"
	"strings"
	"time"
)

// RandStringGenerator - тип для генератора случайных строк.
type RandStringGenerator func(length int) string

// New - создает новый генератор случайных строк на основе заданного набора символов
// Параметры:
// seed - набор символов, из которого будут генерироваться случайные строки
// Возвращается:
// функция-генератор, которая принимает длину строки и возвращает случайную строку.
func New(seed string) RandStringGenerator {
	// Проверяем, что seed не пустой
	if seed == "" {
		panic("seed cannot be empty")
	}

	// Преобразуем строку в массив рун для корректной работы с Unicode символами
	runes := []rune(seed)
	maxLength := len(runes)

	// Инициализируем генератор случайных чисел текущим временем
	rand.Seed(time.Now().UnixNano())

	// Возвращаем замыкание (генератор)
	return func(length int) string {
		// Проверяем корректность длины
		if length < 0 {
			panic("length cannot be negative")
		}

		// Для нулевой длины возвращаем пустую строку
		if length == 0 {
			return ""
		}

		// Используем strings.Builder для эффективной конкатенации строк
		var sb strings.Builder
		sb.Grow(length) // Предварительное выделение памяти

		// Генерируем случайную строку заданной длины
		for i := 0; i < length; i++ {
			sb.WriteRune(runes[rand.Intn(maxLength)])
		}

		return sb.String()
	}
}
