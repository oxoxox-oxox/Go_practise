package config

import (
	"os"
)

type Config struct {
	ServerPort string
	RedisAddr  string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	return Config{
		ServerPort: port,
		RedisAddr:  redisAddr,
	}
}