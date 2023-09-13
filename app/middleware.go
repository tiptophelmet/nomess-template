package app

import (
	mw "github.com/tiptophelmet/nomess-core/middleware"
	"github.com/tiptophelmet/nomess-template/middleware"
)

func defaultMiddleware() []mw.MiddlewareFunc {
	return []mw.MiddlewareFunc{
	}
}

func useMiddleware(middlwr ...mw.MiddlewareFunc) []mw.MiddlewareFunc {
	return append(middlwr, defaultMiddleware()...)
}

func initMiddleware() {
	mw.Register("/register", useMiddleware(
		middleware.WithLocalize,
	))

	mw.Register("/chat", useMiddleware(
		middleware.WithSession,
	))
}
