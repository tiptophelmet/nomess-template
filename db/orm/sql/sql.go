package sql

import (
	"fmt"

	"github.com/tiptophelmet/nomess/logger"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitGormConnection(dialector gorm.Dialector, config *gorm.Config) *gorm.DB {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		logger.Emergency(fmt.Sprintf("could not resolve connect to postgres db: %v", err.Error()))
	}

	prepareConnectionPool()

	return db
}

func prepareConnectionPool() {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Emergency(fmt.Sprintf("could not resolve connect to postgres db: %v", err.Error()))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func Connection() *gorm.DB {
	return db
}
