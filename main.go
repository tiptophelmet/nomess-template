package main

import (
	"fmt"
	"net/http"

	"github.com/tiptophelmet/nomess-template/app"
	"github.com/tiptophelmet/nomess-template/internal/config"
	"github.com/tiptophelmet/nomess-template/internal/logger"
	"github.com/tiptophelmet/nomess-template/internal/router"
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
