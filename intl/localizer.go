package intl

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

var lz *Localizer

func Init(defaultLocale string) *Localizer {
	if lz != nil {
		return lz
	}

	lz = &Localizer{locale: defaultLocale}

	loadLocale()

	return lz
}

func loadLocale() {
	var err error

	lz.localeTree, err = toml.LoadFile(fmt.Sprintf("../locales/%s.toml", lz.locale))

	if err != nil {
		panic(err)
	}
}

type Localizer struct {
	locale     string
	localeTree *toml.Tree
}

func SetLocale(locale string) {
	lz.locale = locale

	loadLocale()
}

func GetLocale() string {
	return lz.locale
}

func Localize(key string) string {
	value := lz.localeTree.Get(key)
	if value != nil {
		if s, ok := value.(string); ok {
			return s
		}
	}
	return ""
}
