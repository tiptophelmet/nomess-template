package app

import "fmt"

func InitApp() {
	initLogger()
	initConfigs()

	initLocales()

	// initSession()

	initMiddleware()
	initPostProcessors()
	initRoutes()

	initDB()
	// initCache()
	// initPubSub()

	fmt.Println("App init OK!")
}
