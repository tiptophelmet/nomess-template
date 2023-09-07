package app

import (
	postproc "github.com/tiptophelmet/nomess-template/internal/postprocessor"
	"github.com/tiptophelmet/nomess-template/postprocessor"
)

func initPostProcessors() {
	postproc.Register("/register", []postproc.PostProcFunc{
		postprocessor.WithLogging,
	})
}
