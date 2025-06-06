package randstring

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	enCharacterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ruCharacterSet = "абвгдеёжхийклмнопрстуфхцчшщьыъэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЬЫЪЭЮЯ"
)

func BenchmarkRandStringGenerator(b *testing.B) {
	generator := New(enCharacterSet)
	for i := 0; i < b.N; i++ {
		generator(16)
	}
}

func TestNew_ValidInput(t *testing.T) {
	generator := New(enCharacterSet)
	result := generator(5)

	assert.Equal(t, 5, len(result), "Длина результата должна быть равна 5")

	for _, char := range result {
		assert.True(t, strings.ContainsRune(enCharacterSet, char),
			"Найден недопустимый символ: %c", char)
	}
}

func TestNew_ZeroLength(t *testing.T) {
	generator := New(enCharacterSet)
	result := generator(0)

	assert.Empty(t, result, "При длине 0 должна возвращаться пустая строка")
}

func TestNew_NegativeLength(t *testing.T) {
	generator := New(enCharacterSet)

	require.Panics(t, func() { generator(-5) },
		"Должен быть вызван panic при отрицательной длине")
}

func TestNew_EmptySeed(t *testing.T) {
	require.Panics(t, func() { New("") },
		"Должен быть вызван panic при пустой строке")
}

func TestNew_UnicodeCharacters(t *testing.T) {
	generator := New(ruCharacterSet)
	result := generator(5)

	assert.Equal(t, 10, len(result), "Длина результата должна быть равна 10")

	for _, char := range result {
		assert.True(t, strings.ContainsRune(ruCharacterSet, char),
			"Найден недопустимый символ: %c", char)
	}
}

func TestNew_MultipleCalls(t *testing.T) {
	generator := New(enCharacterSet)
	results := make(map[string]struct{})

	// Генерируем 100 строк
	for i := 0; i < 100; i++ {
		result := generator(5)
		results[result] = struct{}{}
	}

	assert.GreaterOrEqual(t, len(results), 100,
		"Ожидалось 1000 уникальных строк, получено: %d", len(results))
}

func TestNew_DifferentSeeds(t *testing.T) {
	generator1 := New(enCharacterSet)
	generator2 := New(ruCharacterSet)

	result1 := generator1(5)
	result2 := generator2(5)

	assert.True(t, strings.ContainsAny(result1, enCharacterSet),
		"Результат должен содержать символы из строки '%s': %s", enCharacterSet, result1)

	assert.True(t, strings.ContainsAny(result2, ruCharacterSet),
		"Результат должен содержать символы из строки '%s': %s", ruCharacterSet, result2)
}

func TestNew_SeedUniqueness(t *testing.T) {
	generator1 := New(enCharacterSet)
	generator2 := New(ruCharacterSet)

	result1 := generator1(5)
	result2 := generator2(5)

	assert.NotEqual(t, result1, result2,
		"Результаты должны отличаться при разных исходных строк")
}

func TestNew_Consistency(t *testing.T) {
	generator := New(enCharacterSet)
	result1 := generator(5)
	result2 := generator(5)

	assert.NotEqual(t, result1, result2,
		"Результаты должны быть разными при одинаковых вызовах")
}
