package postprocessor

import (
	"net/http"

	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-core/v5/util"
)

func WithLogging(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	logger.Info("Request to [%s] %s finished", r.Method, util.GetRoutePattern(r))
	
	return w, r
}
