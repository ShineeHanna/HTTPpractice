package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct { //ссылка на бд(структура для создания таблицы)
	gorm.Model        //содержит в себе все необходимые поля для структуры, чтобы самому всё не писать подряд
	Url        string `json:"url"`
	Hash       string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link { //формсирование новой ссылки, когда пришел новый пользователь
	return &Link{
		Url:  url,
		Hash: RandStringRunes(10),
	}
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ")

func RandStringRunes(n int) string { //генерирует случайную строку из рун
	b := make([]rune, n) //сделать массив рун
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))] //берем случайный индекс в letterrunes и записываем в b
	}
	return string(b)
}
