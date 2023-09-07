package middleware

import (
	"net/http"
	"strings"

	"github.com/tiptophelmet/nomess-template/internal/intl"
	"github.com/tiptophelmet/nomess-template/internal/locales"
	"github.com/tiptophelmet/nomess-template/util"
)

func WithLocalize(w http.ResponseWriter, r *http.Request) {
	locale := r.Header.Get("X-Chosen-Language")

	if locales.IsSupportedLocale(locale) || strings.TrimSpace(locale) != "" {
		intl.SetLocale(locale)
	} else if langTags := r.Header.Get("Accept-Language"); langTags != "" {
		intl.SetLocale(util.PickAcceptLang(langTags))
	}
}
