package locales

import "strings"

var SupportedLocales = []string{"en-US"}

func IsSupportedLocale(tag string) bool {
	for _, lang := range SupportedLocales {
		if strings.HasPrefix(tag, lang) {
			return true
		}
	}
	return false
}
