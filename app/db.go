package app

import (
	"fmt"

	"github.com/tiptophelmet/nomess-core/v3/config"
	"github.com/tiptophelmet/nomess-core/v3/db/orm"
	"github.com/tiptophelmet/nomess-core/v3/db/orm/sql"
	"github.com/tiptophelmet/nomess-template/model"
)

func initORM() {
	driver := config.Get("db.orm.driver").Required().Str()
	dsn := config.Get("db.orm.dsn").Required().Str()

	orm.Init(driver, dsn)

	fmt.Println("DB init OK!")
}

func initDB() {
	initORM()
	runMigrations()
}

func runMigrations() {
	sql.Connection().AutoMigrate(&model.Item{})
}
