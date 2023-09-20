package app

import (
	mw "github.com/tiptophelmet/nomess-core/v5/middleware"
	"github.com/tiptophelmet/nomess-template/middleware"
)

func defaultMiddleware() []mw.MiddlewareFunc {
	return []mw.MiddlewareFunc{
		middleware.WithRequestID,
	}
}

func initMiddleware() {
	mw.RegisterMulti([]string{"/item", "/item/{id}"}, []mw.MiddlewareFunc{
		middleware.WithLocalize,
	})

	mw.Default(defaultMiddleware())
}
