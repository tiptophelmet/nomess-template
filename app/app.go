package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/intl"
	"github.com/tiptophelmet/nomess/logger"
	mw "github.com/tiptophelmet/nomess/middleware"
	"github.com/tiptophelmet/nomess/responder"
)

type App struct {
	router *mux.Router
}

func InitApp() *App {
	logger.Init()

	config.Init()

	intl.Init("en-US")

	return &App{router: mux.NewRouter()}
}

// enhanced http handler with middleware support
func (app *App) Handle(pattern string, handlr func(http.ResponseWriter, *http.Request)) {
	app.router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		responder.Init(w, r)

		mw.WithMiddleware(w, r)

		handlr(w, r)
	})
}

// enhanced websocket handler
func (app *App) WebSocket(pattern string, upgrader *websocket.Upgrader, handlr func(*websocket.Conn)) {
	app.router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		mw.WithMiddleware(w, r)

		ws, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			logger.Err(fmt.Sprintf("failed to upgrade HTTP connection for %v with error: %v", pattern, err.Error()))
			return
		}

		handlr(ws)
	})
}
