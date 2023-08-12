package config

var supportedConfigs = map[string]string{
	"mongodb-uri": "NOMESS_MONGODB_URI",
	"mailer":      "NOMESS_MAILER",
}

func getSupportedConfigKeys() []string {
	var keys []string

	for k := range supportedConfigs {
		keys = append(keys, k)
	}

	return keys
}
