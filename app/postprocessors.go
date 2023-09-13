package app

import (
	postproc "github.com/tiptophelmet/nomess-core/postprocessor"
	"github.com/tiptophelmet/nomess-template/postprocessor"
)

func defaultPostProcessors() []postproc.PostProcFunc {
	return []postproc.PostProcFunc{
		postprocessor.WithLogging,
		postprocessor.WithCompression,
		postprocessor.WithStrictTransportSecurity,
		postprocessor.WithContentSecurityPolicy,
	}
}

func usePostProcs(postProcs ...postproc.PostProcFunc) []postproc.PostProcFunc {
	return append(postProcs, defaultPostProcessors()...)
}

func initPostProcessors() {
	postproc.Register("/register", usePostProcs(
		// List specific postprocessors
	))
}
