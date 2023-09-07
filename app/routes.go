package app

import (
	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess-core/router"
	"github.com/tiptophelmet/nomess-template/handler"
)

func initRoutes() {
	router.Init()
	router.Handle("/register", handler.Register).Methods("GET")
	router.WebSocket("/chat", &websocket.Upgrader{}, handler.Chat)
}
