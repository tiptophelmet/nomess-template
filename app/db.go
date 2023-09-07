package app

import (
	"github.com/tiptophelmet/nomess-template/internal/config"
	"github.com/tiptophelmet/nomess-template/internal/db/orm"
)

func initDB() {
	driver := config.Get("db.orm.driver").Required().Str()
	dsn := config.Get("db.orm.dsn").Required().Str()

	orm.Init(driver, dsn)
}
