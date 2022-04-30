package app

import (
	"os"
	"time"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
	Telegram        *TelegramConfig
}

type TelegramConfig struct {
	Token             string
	LongPollerTimeout time.Duration
}

func LoadConfig() *Config {
	c := &Config{
		Port:            "8080",
		ShutdownTimeout: 60 * time.Second,
		Telegram: &TelegramConfig{
			LongPollerTimeout: 60 * time.Second,
		},
	}
	if port, ok := os.LookupEnv("PORT"); ok {
		c.Port = port
	}
	if token, ok := os.LookupEnv("TELEGRAM_TOKEN"); ok {
		c.Telegram.Token = token
	}
	return c
}
