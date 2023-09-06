package app

import (
	postproc "github.com/tiptophelmet/nomess/internal/postprocessor"
	"github.com/tiptophelmet/nomess/postprocessor"
)

func initPostProcessors() {
	postproc.Register("/register", []postproc.PostProcFunc{
		postprocessor.WithLogging,
	})
}
