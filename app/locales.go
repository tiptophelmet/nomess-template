package app

import (
	"github.com/tiptophelmet/nomess-core/intl"
	"github.com/tiptophelmet/nomess-core/locales"
)

func initLocales() {
	locales.Register([]string{"en-US"})

	intl.Init("en-US")
}
