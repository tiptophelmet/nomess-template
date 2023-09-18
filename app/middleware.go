package app

import (
	mw "github.com/tiptophelmet/nomess-core/v3/middleware"
	"github.com/tiptophelmet/nomess-template/middleware"
)

func defaultMiddleware() []mw.MiddlewareFunc {
	return []mw.MiddlewareFunc{
		middleware.WithRequestID,
	}
}

func useMiddleware(middlwr ...mw.MiddlewareFunc) []mw.MiddlewareFunc {
	return append(middlwr, defaultMiddleware()...)
}

func initMiddleware() {
	mw.Register("/item", useMiddleware(
		middleware.WithLocalize,
	))
}
