package app

import (
	"github.com/tiptophelmet/nomess/internal/intl"
	"github.com/tiptophelmet/nomess/internal/locales"
)

func initLocales() {
	locales.Register([]string{"en-US"})

	intl.Init("en-US")
}
