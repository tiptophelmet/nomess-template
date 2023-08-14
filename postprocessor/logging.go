package postprocessor

import (
	"fmt"
	"net/http"

	"github.com/tiptophelmet/nomess/logger"
)

func WithLogging(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("Request %v finished with response %v", r, w))
}
