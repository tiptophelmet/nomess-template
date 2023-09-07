package app

func InitApp() {
	initLogger()
	initConfigs()

	initLocales()

	initSession()

	initMiddleware()
	initPostProcessors()
	initRoutes()

	initDB()
	initCache()
	initPubSub()
}
