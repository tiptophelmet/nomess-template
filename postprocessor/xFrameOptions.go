package postprocessor

import (
	"net/http"
)

func WithXFrameOptions(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	w.Header().Add("X-Frame-Options", "deny")

	return w, r
}
