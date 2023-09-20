package app

import (
	postproc "github.com/tiptophelmet/nomess-core/v5/postprocessor"
	"github.com/tiptophelmet/nomess-template/postprocessor"
)

func defaultPostProcessors() []postproc.PostProcFunc {
	return []postproc.PostProcFunc{
		postprocessor.WithCompression,
		postprocessor.WithStrictTransportSecurity,
		postprocessor.WithContentSecurityPolicy,
	}
}

func initPostProcessors() {
	postproc.RegisterMulti([]string{"/item", "/item/{id}"}, []postproc.PostProcFunc{
		postprocessor.WithLogging,
	})

	postproc.Default(defaultPostProcessors())
}
