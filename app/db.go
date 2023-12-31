package app

import (
	"github.com/tiptophelmet/nomess-core/v5/config"
	"github.com/tiptophelmet/nomess-core/v5/db/orm"
	"github.com/tiptophelmet/nomess-core/v5/db/orm/sql"
	"github.com/tiptophelmet/nomess-template/model"
)

func initORM() {
	driver := config.Get("db.orm.driver").Required().Str()
	dsn := config.Get("db.orm.dsn").Required().Str()

	orm.Init(driver, dsn)
}

func initDB() {
	initORM()
	runMigrations()
}

func runMigrations() {
	sql.Connection().AutoMigrate(&model.Item{})
}
