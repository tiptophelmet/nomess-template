package main

import (
	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess/app"
	"github.com/tiptophelmet/nomess/handler"
)

func main() {
	app := app.InitApp()

	app.Handle("/api/auth/register", handler.Register)
	app.WebSocket("/chat", &websocket.Upgrader{}, handler.Chat)
}
