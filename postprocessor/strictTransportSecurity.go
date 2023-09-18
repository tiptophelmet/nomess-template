package postprocessor

import (
	"fmt"
	"net/http"

	"github.com/tiptophelmet/nomess-core/v3/config"
)

func WithStrictTransportSecurity(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	maxAgeSeconds := config.Get("strict-transport-security.max-age").Int()

	headerVal := fmt.Sprintf("max-age=%d; includeSubDomains; preload", maxAgeSeconds)
	w.Header().Add("Strict-Transport-Security", headerVal)

	return w, r
}
