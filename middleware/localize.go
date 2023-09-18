package middleware

import (
	"net/http"
	"strings"

	"github.com/tiptophelmet/nomess-core/v3/intl"
	"github.com/tiptophelmet/nomess-core/v3/locales"
	"github.com/tiptophelmet/nomess-template/util"
)

func WithLocalize(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	locale := r.Header.Get("X-Chosen-Language")

	if locales.IsSupportedLocale(locale) || strings.TrimSpace(locale) != "" {
		intl.SetLocale(locale)
	} else if langTags := r.Header.Get("Accept-Language"); langTags != "" {
		intl.SetLocale(util.PickAcceptLang(langTags))
	}

	return w, r
}
