package repo

import (
	"github.com/tiptophelmet/nomess-core/v3/db/orm/sql"
	"github.com/tiptophelmet/nomess-template/model"
	"gorm.io/gorm"
)

type Item struct {
	db *gorm.DB
}

func InitItemRepo() *Item {
	return &Item{db: sql.Connection()}
}

func (repo *Item) Save(item *model.Item) (*model.Item, error) {
	result := repo.db.Create(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}
