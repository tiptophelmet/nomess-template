package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/intl"
	"github.com/tiptophelmet/nomess/logger"
	mw "github.com/tiptophelmet/nomess/middleware"
	"github.com/tiptophelmet/nomess/responder"
)

func initApp() {
	logger.Init()

	config.Init()

	intl.Init("en-US")
}

// enhanced http handler with middleware support
func Handle(pattern string, handlr func(http.ResponseWriter, *http.Request)) {
	initApp()

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		responder.Init(w, r)

		mw.WithMiddleware(w, r)

		handlr(w, r)
	})
}

// enhanced websocket handler with middleware support
func WebSocket(pattern string, upgrader *websocket.Upgrader, handlr func(*websocket.Conn)) {
	initApp()

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		mw.WithMiddleware(w, r)

		ws, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			logger.Err(fmt.Sprintf("failed to upgrade HTTP connection for %v with error: %v", pattern, err.Error()))
			return
		}

		handlr(ws)
	})
}
