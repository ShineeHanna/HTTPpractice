package link

import (
	"pack/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct { //база данных полученных|созданных ссылок link
	Database *db.Db
}

func NewLinkRepositiry(database *db.Db) *LinkRepository { //формирование новой ссылки на базу данных
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link) //запись созданной ссылки в БД
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) { //взять из базы запись
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash) //взять значение из базы данных и вместо ? вставиться hash
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) { //одновить запись в таблице
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error { //одновить запись в таблице
	result := repo.Database.DB.Delete(&Link{}, id) //удалить запись по id
	if result.Error != nil {
		return result.Error
	}
	return nil
}
