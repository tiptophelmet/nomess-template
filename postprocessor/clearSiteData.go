package postprocessor

import (
	"fmt"
	"net/http"
	"strings"
)

func clearSiteData(directives ...string) string {
	var directiveStr string
	for _, d := range directives {
		directiveStr += fmt.Sprintf("\"%s\", ", d)
	}

	return strings.TrimRight(directiveStr, ", ")
}

func WithClearSiteData(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	w.Header().Add("Clear-Site-Data", clearSiteData("cache", "cookies", "storage", "executionContexts"))
	
	return w, r
}
