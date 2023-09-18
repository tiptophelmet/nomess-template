package postprocessor

import (
	"net/http"

	"github.com/tiptophelmet/nomess-core/v3/logger"
)

func WithLogging(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	logger.Info("Request %v finished with response %v", r, w)

	return w, r
}
