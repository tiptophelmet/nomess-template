package app

import (
	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess-template/handler"
	"github.com/tiptophelmet/nomess-template/internal/router"
)

func initRoutes() {
	router.Init()
	router.Handle("/register", handler.Register).Methods("GET")
	router.WebSocket("/chat", &websocket.Upgrader{}, handler.Chat)
}
