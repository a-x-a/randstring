# RANDSTRING

[![golangci-lint](https://github.com/a-x-a/randstring/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/a-x-a/randstring/actions/workflows/golangci-lint.yml)
[![Test and coverage](https://github.com/a-x-a/randstring/actions/workflows/codecovtest.yml/badge.svg)](https://github.com/a-x-a/randstring/actions/workflows/codecovtest.yml)

## Генератор случайных строк на основе заданного набора символов

### Описание

Генератор случайных строк, на основе заданного набора символов. В качестве набора символов могут использоваться символы из различных языков. Генератор возвращает функцию, которая используется для генерации строки заданной длины.

### Использование

```go
package main

import (
	"fmt"

	"github.com/a-x-a/randstring"
)

func main() {
	// Создаем генератор с буквами и цифрами
	generator := randstring.New("abcdefghijklmnopqrstuvwxyz0123456789")

	// Генерируем пароль длиной 8 символов
	password := generator(8)
	fmt.Println(password)

	// Генерируем тестовые данные длиной 10 символов
	testData := generator(10)
	fmt.Println(testData)
}
```

### Применение

- Генерация паролей
- Создание тестовых данных
- Генерация случайных идентификаторов
- Создание случайных строк для валидации

### Рекомендации

- Для криптографической безопасности используйте специализированные библиотеки
- Начальный набор символов должен содержать достаточное количество различных символов для получения хорошего распределения
- При использовании в многопоточной среде рекомендуется создавать отдельный генератор для каждого потока
