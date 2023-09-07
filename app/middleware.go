package app

import (
	mw "github.com/tiptophelmet/nomess-core/middleware"
	"github.com/tiptophelmet/nomess-template/middleware"
)

func initMiddleware() {
	mw.Register("/register", []mw.MiddlewareFunc{
		middleware.WithLocalize,
	})

	mw.Register("/chat", []mw.MiddlewareFunc{
		middleware.WithSession,
	})
}
