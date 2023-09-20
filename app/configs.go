package app

import (
	"embed"

	"github.com/tiptophelmet/nomess-core/v5/config"
)

var configs = map[string]string{
	"port": "NOMESS_PORT",
	// "session.jwt.expiration.time":       "NOMESS_SESSION_JWT_EXPIRATION_TIME",
	// "session.jwt.expiration.window":     "NOMESS_SESSION_JWT_EXPIRATION_WINDOW",
	// "session.jwt.secret":                "NOMESS_SESSION_JWT_SECRET",
	"db.orm.driver": "NOMESS_DB_ORM_DRIVER",
	"db.orm.dsn":    "NOMESS_DB_ORM_DSN",
	// "cache.driver":                      "NOMESS_CACHE_DRIVER",
	// "cache.url":                         "NOMESS_CACHE_URL",
	// "pubsub.driver":                     "NOMESS_PUBSUB_DRIVER",
	// "pubsub.url":                        "NOMESS_PUBSUB_URL",
	"strict-transport-security.max-age": "NOMESS_STRICT_TRANSPORT_SECURITY_MAX_AGE",
}

//go:embed config.toml
var fallbackFile embed.FS

func initConfigs() {
	config.Register(configs)
	config.Init(fallbackFile)
}
