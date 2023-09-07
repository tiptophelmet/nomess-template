package repo

import (
	"github.com/tiptophelmet/nomess/internal/db/orm/sql"
	"github.com/tiptophelmet/nomess/model"
	"gorm.io/gorm"
)

type UserVerification struct {
	db *gorm.DB
}

func InitUserVerificationRepo() *UserVerification {
	return &UserVerification{db: sql.Connection()}
}

func (repo *UserVerification) Save(userVerification *model.UserVerification) error {
	result := repo.db.Create(&userVerification)
	return result.Error
}
