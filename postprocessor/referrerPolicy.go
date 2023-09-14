package postprocessor

import (
	"net/http"
)

func WithReferrerPolicy(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	w.Header().Add("Referrer-Policy", "origin-when-cross-origin")

	return w, r
}
