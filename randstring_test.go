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
	generator, _ := NewGenerator(enCharacterSet)
	for i := 0; i < b.N; i++ {
		generator(16)
	}
}

func TestNew_ValidInput(t *testing.T) {
	generator,err := NewGenerator(enCharacterSet)
	assert.NoError(t, err, "Должен быть успешный результат")

	result, err := generator(5)
	assert.NoError(t, err, "Должен быть успешный результат")
	assert.Equal(t, 5, len(result), "Длина результата должна быть равна 5")

	for _, char := range result {
		assert.True(t, strings.ContainsRune(enCharacterSet, char),
			"Найден недопустимый символ: %c", char)
	}
}

func TestNew_ZeroLength(t *testing.T) {
	generator, err := NewGenerator(enCharacterSet)
	assert.NoError(t, err, "Должен быть успешный результат")

	result,err := generator(0)
	assert.NoError(t, err, "Должен быть успешный результат")
	assert.Empty(t, result, "При длине 0 должна возвращаться пустая строка")
}

func TestNew_NegativeLength(t *testing.T) {
	generator, err := NewGenerator(enCharacterSet)
	assert.NoError(t, err, "Должен быть успешный результат")

	_, err = generator(-5)
	assert.EqualError(t, err, ErrNegativeLength.Error())
}

func TestNew_EmptySeed(t *testing.T) {
	_, err := NewGenerator("")
	assert.EqualError(t, err, ErrEmptySeed.Error())
}

func TestNew_UnicodeCharacters(t *testing.T) {
	generator, err := NewGenerator(ruCharacterSet)
	require.NoError(t, err, "Должен быть успешный результат")

	result, err := generator(5)
	assert.NoError(t, err, "Должен быть успешный результат")
	assert.Equal(t, 10, len(result), "Длина результата должна быть равна 10")

	for _, char := range result {
		assert.True(t, strings.ContainsRune(ruCharacterSet, char),
			"Найден недопустимый символ: %c", char)
	}
}

func TestNew_MultipleCalls(t *testing.T) {
	generator, err := NewGenerator(enCharacterSet)
	require.NoError(t, err, "Должен быть успешный результат")

	results := make(map[string]struct{})

	// Генерируем 100 строк
	for i := 0; i < 100; i++ {
		result, err := generator(5)
		require.NoError(t, err, "Должен быть успешный результат")
		results[result] = struct{}{}
	}

	assert.GreaterOrEqual(t, len(results), 100,
		"Ожидалось 1000 уникальных строк, получено: %d", len(results))
}

func TestNew_DifferentSeeds(t *testing.T) {
	generator1, err := NewGenerator(enCharacterSet)
	require.NoError(t, err, "Должен быть успешный результат")
	
	generator2, err := NewGenerator(ruCharacterSet)
	require.NoError(t, err, "Должен быть успешный результат")

	result1, err := generator1(5)
	require.NoError(t, err, "Должен быть успешный результат")

	result2, err := generator2(5)
	require.NoError(t, err, "Должен быть успешный результат")

	assert.True(t, strings.ContainsAny(result1, enCharacterSet),
		"Результат должен содержать символы из строки '%s': %s", enCharacterSet, result1)

	assert.True(t, strings.ContainsAny(result2, ruCharacterSet),
		"Результат должен содержать символы из строки '%s': %s", ruCharacterSet, result2)
}

func TestNew_SeedUniqueness(t *testing.T) {
	generator1, err := NewGenerator(enCharacterSet)
	require.NoError(t, err, "Должен быть успешный результат")
	
	generator2, err := NewGenerator(ruCharacterSet)
	require.NoError(t, err, "Должен быть успешный результат")

	result1, err := generator1(5)
	require.NoError(t, err, "Должен быть успешный результат")

	result2, err := generator2(5)
	require.NoError(t, err, "Должен быть успешный результат")

	assert.NotEqual(t, result1, result2,
		"Результаты должны отличаться при разных исходных строк")
}

func TestNew_Consistency(t *testing.T) {
	generator, err := NewGenerator(enCharacterSet)
	require.NoError(t, err, "Должен быть успешный результат")

	result1, err := generator(5)
	require.NoError(t, err, "Должен быть успешный результат")

	result2, err := generator(5)
	require.NoError(t, err, "Должен быть успешный результат")

	assert.NotEqual(t, result1, result2,
		"Результаты должны быть разными при одинаковых вызовах")
}
