package orm

import (
	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/db/orm/sql"
	"github.com/tiptophelmet/nomess/logger"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Init() {
	driverConfig := config.Get("db.orm.driver").Required().Str()
	dsn := config.Get("db.orm.dsn").Required().Str()

	var dialector gorm.Dialector

	switch driverConfig {
	case "mysql":
	case "tidb":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	case "sqlserver":
		dialector = sqlserver.Open(dsn)
	case "clickhouse":
		dialector = clickhouse.Open(dsn)
	default:
		logger.Panic("unsupported db.orm.driver: %v", driverConfig)
	}

	sql.InitGormConnection(dialector, &gorm.Config{TranslateError: true})
}
