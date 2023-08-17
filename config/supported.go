package config

var supportedConfigs = map[string]string{
	"port":                          "NOMESS_PORT",
	"mongodb-uri":                   "NOMESS_MONGODB_URI",
	"mailer":                        "NOMESS_MAILER",
	"session.jwt.expiration.time":   "NOMESS_SESSION_JWT_EXPIRATION_TIME",
	"session.jwt.expiration.window": "NOMESS_SESSION_JWT_EXPIRATION_WINDOW",
	"session.jwt.secret":            "NOMESS_SESSION_JWT_SECRET",
	"db.orm.driver":                 "NOMESS_DB_ORM_DRIVER",
	"db.orm.dsn":                    "NOMESS_DB_ORM_DSN",
}

func getSupportedConfigKeys() []string {
	var keys []string

	for k := range supportedConfigs {
		keys = append(keys, k)
	}

	return keys
}
