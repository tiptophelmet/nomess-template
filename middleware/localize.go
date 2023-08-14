package middleware

import (
	"net/http"
	"strings"

	"github.com/tiptophelmet/nomess/intl"
	"github.com/tiptophelmet/nomess/locales"
	"github.com/tiptophelmet/nomess/util"
)

func WithLocalize(w http.ResponseWriter, r *http.Request) {
	locale := r.Header.Get("X-Chosen-Language")

	if locales.IsSupportedLocale(locale) || strings.TrimSpace(locale) != "" {
		intl.SetLocale(locale)
	} else if langTags := r.Header.Get("Accept-Language"); langTags != "" {
		intl.SetLocale(util.PickAcceptLang(langTags))
	}
}
