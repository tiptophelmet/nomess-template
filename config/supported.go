package config

var supportedConfigs = map[string]string{
	"port":                          "NOMESS_PORT",
	"mongodb.uri":                   "NOMESS_MONGODB_URI",
	"mail.driver":                   "NOMESS_MAIL_DRIVER",
	"session.jwt.expiration.time":   "NOMESS_SESSION_JWT_EXPIRATION_TIME",
	"session.jwt.expiration.window": "NOMESS_SESSION_JWT_EXPIRATION_WINDOW",
	"session.jwt.secret":            "NOMESS_SESSION_JWT_SECRET",
	"db.orm.driver":                 "NOMESS_DB_ORM_DRIVER",
	"db.orm.dsn":                    "NOMESS_DB_ORM_DSN",
	"cache.driver":                  "NOMESS_CACHE_DRIVER",
	"cache.url":                     "NOMESS_CACHE_URL",
	"pubsub.driver":                 "NOMESS_PUBSUB_DRIVER",
	"pubsub.url":                    "NOMESS_PUBSUB_URL",
}

func getSupportedConfigKeys() []string {
	var keys []string

	for k := range supportedConfigs {
		keys = append(keys, k)
	}

	return keys
}
