package response

import "github.com/tiptophelmet/nomess-core/v5/intl"

type badRequest struct {
	Message string
	Text    string
}

func BadRequest() *badRequest {
	return &badRequest{
		Message: intl.Localize("bad_request.message"),
		Text:    intl.Localize("bad_request.text"),
	}
}
