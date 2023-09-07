package repo

import (
	"github.com/tiptophelmet/nomess-core/db/orm/sql"
	"github.com/tiptophelmet/nomess-template/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func InitUserRepo() *User {
	return &User{db: sql.Connection()}
}

func (repo *User) Save(user *model.User) error {
	result := repo.db.Create(&user)
	return result.Error
}
