package app

import (
	"github.com/tiptophelmet/nomess-template/internal/intl"
	"github.com/tiptophelmet/nomess-template/internal/locales"
)

func initLocales() {
	locales.Register([]string{"en-US"})

	intl.Init("en-US")
}
