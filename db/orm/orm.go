package orm

import (
	"fmt"

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

func Init() *gorm.DB {
	driverConfig, err := config.Str("db.orm.driver")
	if err != nil {
		logger.Alert("could not resolve db.orm.driver")
	}

	dsn, err := config.Str("db.orm.dsn")
	if err != nil {
		logger.Emergency("could not resolve db.orm.dsn")
	}

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
		logger.Emergency(fmt.Sprintf("unsupported db.orm.driver: %v", driverConfig))
		return nil
	}

	return sql.InitGormConnection(dialector, &gorm.Config{TranslateError: true})
}
