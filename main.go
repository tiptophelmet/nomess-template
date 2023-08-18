package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess/app"
	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/handler"
	"github.com/tiptophelmet/nomess/logger"
	"github.com/tiptophelmet/nomess/router"
)

func main() {
	app.Init()

	prepareRoutes()
	startServer()
}

func prepareRoutes() {
	router.Handle("/register", handler.Register).Methods("GET")
	router.WebSocket("/chat", &websocket.Upgrader{}, handler.Chat)
}

func startServer() {
	port := config.Get("port").Required().Str()
	prefixedPort := fmt.Sprintf(":%v", port)

	err := http.ListenAndServe(prefixedPort, router.GetMuxRouter())
	if err != nil {
		logger.Emergency(fmt.Sprintf("failed to start http server: %v", err.Error()))
	} else {
		logger.Debug("server started at port %v")
	}
}
