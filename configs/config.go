package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load() //получаем зависимости из файла .env
	if err != nil {
		log.Println("Erorr loading .env fule, using default config")
	}
	return &Config{ //при успешном получении присваиваем значение в наши поля
		Db: DbConfig{
			Dsn: os.Getenv("DSN"), //получаем значение по названию DSN
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"), //получаем значение по названию TOKEN
		},
	}
}
