package config

var supportedConfigs = map[string]string{
	"port":                          "NOMESS_PORT",
	"mongodb-uri":                   "NOMESS_MONGODB_URI",
	"mailer":                        "NOMESS_MAILER",
	"session.jwt.expiration.time":   "NOMESS_SESSION_JWT_EXPIRATION_TIME",
	"session.jwt.expiration.window": "NOMESS_SESSION_JWT_EXPIRATION_WINDOW",
	"session.jwt.secret":            "NOMESS_SESSION_JWT_SECRET",
}

func getSupportedConfigKeys() []string {
	var keys []string

	for k := range supportedConfigs {
		keys = append(keys, k)
	}

	return keys
}
