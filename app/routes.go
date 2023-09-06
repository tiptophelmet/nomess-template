package app

import (
	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess/handler"
	"github.com/tiptophelmet/nomess/internal/router"
)

func initRoutes() {
	router.Init()
	router.Handle("/register", handler.Register).Methods("GET")
	router.WebSocket("/chat", &websocket.Upgrader{}, handler.Chat)
}
