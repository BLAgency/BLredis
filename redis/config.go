package redis

import "time"

type Config struct {
	Addr         string        // Адрес Redis сервера
	Password     string        // Пароль (если требуется)
	DB           int           // Номер базы данных
	PoolSize     int           // Размер пула соединений
	MinIdleConns int           // Минимальное количество idle соединений
	MaxConnAge   time.Duration // Максимальный возраст соединения
	PoolTimeout  time.Duration // Таймаут пула
	IdleTimeout  time.Duration // Таймаут idle соединений
}

func DefaultConfig() *Config {
	return &Config{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
		MaxConnAge:   30 * time.Minute,
		PoolTimeout:  4 * time.Second,
		IdleTimeout:  5 * time.Minute,
	}
}
