package postprocessor

import (
	"net/http"

	"github.com/tiptophelmet/cspolicy"
	"github.com/tiptophelmet/cspolicy/directives"
	"github.com/tiptophelmet/cspolicy/src"
)

func WithContentSecurityPolicy(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	csp := cspolicy.Build(
		directives.DefaultSrc(src.None()),
		directives.BaseURI(src.Self(), src.Host("*.example.com")),
		directives.ChildSrc(
			src.Host("cdn.example.com/assets"),
			src.Host("resources.example.com/artifacts"),
		),
		directives.ConnectSrc(
			src.Host("uploads.example.com"),
			src.Host("status.example.com"),
			src.Host("api.example.com"),
		),
		directives.FrameSrc(
			src.Host("notes.example.com"),
			src.Host("viewbox.example.com"),
		),
		directives.ImgSrc(
			src.Self(),
			src.Scheme("data:"),
			src.Host("media.example.com"),
			src.Host("avatars.example.com"),
		),
		directives.UpgradeInsecureRequests(),
	)

	w.Header().Add("Content-Security-Policy", csp)
	return w, r
}
