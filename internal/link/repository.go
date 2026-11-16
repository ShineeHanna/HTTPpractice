package link

import "pack/pkg/db"

type LinkRepository struct { //зависимость от базы данных
	Database *db.Db
}

func NewLinkRepositiry(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) {

}
