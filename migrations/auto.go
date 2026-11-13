package main //миграция базы данных в саму базу posgre

import (
	"os"
	"pack/internal/link"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env") //читаем env файл
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{}) //создаём подключение
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{}) //вызываем миграцию
}
