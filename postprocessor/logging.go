package postprocessor

import (
	"net/http"

	"github.com/tiptophelmet/nomess/logger"
)

func WithLogging(w http.ResponseWriter, r *http.Request) {
	logger.Info("Request %v finished with response %v", r, w)
}
