package app

import (
	mw "github.com/tiptophelmet/nomess/internal/middleware"
	"github.com/tiptophelmet/nomess/middleware"
)

func initMiddleware() {
	mw.Register("/register", []mw.MiddlewareFunc{
		middleware.WithLocalize,
	})

	mw.Register("/chat", []mw.MiddlewareFunc{
		middleware.WithSession,
	})
}
