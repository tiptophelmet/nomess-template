package repo

import (
	"github.com/tiptophelmet/nomess-core/v5/db/orm/sql"
	"github.com/tiptophelmet/nomess-template/model"
	"gorm.io/gorm"
)

type Item struct {
	db *gorm.DB
}

func InitItemRepo() *Item {
	return &Item{db: sql.Connection()}
}

func (repo *Item) Create(item *model.Item) (*model.Item, error) {
	result := repo.db.Create(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (repo *Item) Save(item *model.Item) (*model.Item, error) {
	result := repo.db.Save(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (repo *Item) Get(ID int) (*model.Item, error) {
	item := &model.Item{}

	result := repo.db.First(item, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}

func (repo *Item) GetAll() ([]*model.Item, error) {
	var items []*model.Item

	result := repo.db.Find(&items)

	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (repo *Item) Delete(ID int) error {
	result := repo.db.Delete(&model.Item{}, ID)

	return result.Error
}
