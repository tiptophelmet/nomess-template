package response

import "github.com/tiptophelmet/nomess-core/v5/intl"

type internalServerError struct {
	Message string
	Text    string
}

func InternalServerError() *internalServerError {
	return &internalServerError{
		Message: intl.Localize("internal_server_error.message"),
		Text:    intl.Localize("internal_server_error.text"),
	}
}
