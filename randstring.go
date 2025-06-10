package randstring

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var ErrEmptySeed = fmt.Errorf("seed cannot be empty")
var ErrNegativeLength = fmt.Errorf("length cannot be negative")

// Generator - тип для генератора случайных строк.
type Generator func(length int) (string, error)

// NewGenerator - создает новый генератор случайных строк на основе заданного набора символов
// Параметры:
// seed - набор символов, из которого будут генерироваться случайные строки
// Возвращается:
// функция-генератор, которая принимает длину строки и возвращает случайную строку.
// В случае если передана отрицательная длина, возвращает пустую строку и ошибку ErrNegativeLength.
// В случае если seed пустой, ошибку ErrEmptySeed.
func NewGenerator(seed string) (Generator, error) {
	// Проверяем, что seed не пустой
	if seed == "" {
		return nil, ErrEmptySeed
	}

	// Преобразуем строку в массив рун для корректной работы с Unicode символами
	runes := []rune(seed)
	maxLength := len(runes)

	// Инициализируем генератор случайных чисел текущим временем
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Возвращаем замыкание (генератор)
	return func(length int) (string, error) {
		// Проверяем корректность длины
		if length < 0 {
			return "", ErrNegativeLength
		}

		// Используем strings.Builder для эффективной конкатенации строк
		var sb strings.Builder
		sb.Grow(length) // Предварительное выделение памяти

		// Генерируем случайную строку заданной длины
		for _ = range length {
			sb.WriteRune(runes[r.Intn(maxLength)])
		}

		return sb.String(), nil
	}, nil
}
