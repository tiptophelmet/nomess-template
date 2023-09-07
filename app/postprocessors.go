package app

import (
	postproc "github.com/tiptophelmet/nomess-core/postprocessor"
	"github.com/tiptophelmet/nomess-template/postprocessor"
)

func initPostProcessors() {
	postproc.Register("/register", []postproc.PostProcFunc{
		postprocessor.WithLogging,
	})
}
