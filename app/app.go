package app

func InitApp() {
	initLogger()
	initConfigs()
	initLocale()
	initMiddleware()
	initPostProcessors()
	initRoutes()
}
