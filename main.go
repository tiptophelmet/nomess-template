package main

import (
	"fmt"
	"net/http"

	"github.com/tiptophelmet/nomess-core/config"
	"github.com/tiptophelmet/nomess-core/logger"
	"github.com/tiptophelmet/nomess-core/router"
	"github.com/tiptophelmet/nomess-template/app"
)

func main() {
	app.InitApp()

	startServer()
}

func startServer() {
	port := config.Get("port").Required().Str()
	prefixedPort := fmt.Sprintf(":%v", port)

	err := http.ListenAndServe(prefixedPort, router.GetMuxRouter())
	if err != nil {
		logger.Panic("failed to start http server: %v", err.Error())
	} else {
		logger.Debug("server started at port %v", port)
	}
}
