package app

import (
	"embed"

	"github.com/tiptophelmet/nomess-core/v5/intl"
	"github.com/tiptophelmet/nomess-core/v5/locales"
)

//go:embed locales
var localesDir embed.FS

func initLocales() {
	locales.Register([]string{"en-US"})

	intl.Init("en-US", localesDir)
}
