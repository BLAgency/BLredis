# BLredis

BLredis - это пакет для работы с Redis в Go. Он предоставляет удобные функции для подключения, операций с ключами, получения данных, проверки TTL, подсчёта записей и других операций.

## Особенности

- Простое подключение к Redis
- Полный набор операций с ключами (Get, Set, Del, Exists, TTL, Expire)
- Поиск ключей по паттерну (Keys, Scan)
- Подсчёт записей с ключом
- Настраиваемая конфигурация
- Встроенное тестирование

## Установка

```bash
go get github.com/BLAgency/BLredis
```

## Быстрый старт

### Базовое подключение

```go
package main

import (
    "github.com/BLAgency/BLredis/redis"
)

func main() {
    // Инициализация с адресом по умолчанию (localhost:6379)
    client := redis.NewClient()

    // Проверка подключения
    pong, err := client.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println(pong) // PONG
}
```

### Подключение с кастомной конфигурацией

```go
config := &redis.Config{
    Addr:     "localhost:6379",
    Password: "", // если требуется
    DB:       0,
    PoolSize: 10,
}

client := redis.NewClientWithConfig(config)
```

### Основные операции

```go
// Установка значения
err := client.Set("key", "value")
if err != nil {
    // обработка ошибки
}

// Получение значения
value, err := client.Get("key")
if err != nil {
    // обработка ошибки
}
fmt.Println(value) // "value"

// Проверка TTL
ttl, err := client.TTL("key")
if err != nil {
    // обработка ошибки
}
fmt.Println(ttl) // время до истечения в секундах

// Поиск ключей по паттерну
keys, err := client.Keys("prefix:*")
if err != nil {
    // обработка ошибки
}
fmt.Println(keys) // ["prefix:1", "prefix:2", ...]

// Подсчёт записей с ключом (по паттерну)
count, err := client.CountKeys("prefix:*")
if err != nil {
    // обработка ошибки
}
fmt.Println(count) // количество ключей

// Удаление ключа
err = client.Del("key")
if err != nil {
    // обработка ошибки
}

// Проверка существования ключа
exists, err := client.Exists("key")
if err != nil {
    // обработка ошибки
}
fmt.Println(exists) // true или false

// Установка TTL
err = client.Expire("key", 3600) // 1 час
if err != nil {
    // обработка ошибки
}
```

### Сканирование ключей

```go
// Использование Scan для больших наборов данных
cursor := uint64(0)
keys, newCursor, err := client.Scan(cursor, "prefix:*", 10)
if err != nil {
    // обработка ошибки
}
fmt.Println(keys) // ["prefix:1", "prefix:2", ...]

// Продолжение сканирования
for newCursor != 0 {
    keys, newCursor, err = client.Scan(newCursor, "prefix:*", 10)
    if err != nil {
        break
    }
    fmt.Println(keys)
}
```

## API

### Подключение

- `NewClient() *RedisClient` - Создание клиента с конфигом по умолчанию
- `NewClientWithConfig(config *Config) *RedisClient` - Создание клиента с кастомным конфигом

### Операции с ключами

- `Set(key, value string) error` - Установка значения ключа
- `Get(key string) (string, error)` - Получение значения ключа
- `Del(key string) error` - Удаление ключа
- `Exists(key string) (bool, error)` - Проверка существования ключа
- `TTL(key string) (time.Duration, error)` - Получение TTL ключа
- `Expire(key string, ttl int) error` - Установка TTL для ключа
- `Keys(pattern string) ([]string, error)` - Поиск ключей по паттерну
- `Scan(cursor uint64, match string, count int64) ([]string, uint64, error)` - Сканирование ключей
- `CountKeys(pattern string) (int64, error)` - Подсчёт ключей по паттерну

### Конфигурация

```go
type Config struct {
    Addr     string // Адрес Redis сервера
    Password string // Пароль (если требуется)
    DB       int    // Номер базы данных
    PoolSize int    // Размер пула соединений
}
```

## Тестирование

```bash
go test ./redis
```

Для тестов требуется запущенный Redis сервер. Установите переменные окружения:

- `TEST_REDIS_ADDR` - адрес Redis сервера (по умолчанию localhost:6379)
- `TEST_REDIS_PASSWORD` - пароль (если требуется)

## Лицензия

[Укажите лицензию]
