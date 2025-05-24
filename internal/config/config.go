package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config Настройки приложения
type Config struct {
	TelegramAPIKey string
	DatabaseURL    string
	BotToken       string `env:"BOT_TOKEN,required"`
	ChannelID      int64  `env:"CHANNEL_ID,required"`
	DBUser         string `env:"DB_USER" envDefault:"postgres"`
	DBPassword     string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName         string `env:"DB_NAME" envDefault:"adtime-telegram-bot"`
	DBHost         string `env:"DB_HOST" envDefault:"localhost"`
	DBPort         string `env:"DB_PORT" envDefault:"5432"`
	RedisAddr      string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	RedisPassword  string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB        int    `env:"REDIS_DB" envDefault:"0"`
}

// LoadConfig Загрузка из .env
func LoadConfig() *Config {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := &Config{}

	cfg.DatabaseURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	cfg.TelegramAPIKey = os.Getenv("TELEGRAM_API_KEY")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")

	return cfg
}
