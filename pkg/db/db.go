package db

import (
	"pack/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb(conf *configs.Config) *Db { //в этой структуре будет храниться база данных
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}

}

type Db struct {
	*gorm.DB
}
