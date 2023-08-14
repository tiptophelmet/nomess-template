package app

import (
	"net/http"

	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/intl"
	"github.com/tiptophelmet/nomess/logger"
	mw "github.com/tiptophelmet/nomess/middleware"
	"github.com/tiptophelmet/nomess/responder"
)

// enhanced http handler with middleware support
func Handle(pattern string, handlr func(w http.ResponseWriter, r *http.Request)) {
	logger.Init()

	config.Init()

	intl.Init("en-US")

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		mw.WithMiddleware(w, r, handlr)

		responder.Init(w, r)
	})
}
