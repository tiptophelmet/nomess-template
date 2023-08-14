package response

import "github.com/tiptophelmet/nomess/intl"

type internalServerError struct {
	Message string
	Text    string
}

func InternalServerErrorResponse() *internalServerError {
	return &internalServerError{
		Message: intl.Localize("internal_server_error.message"),
		Text:    intl.Localize("internal_server_error.text"),
	}
}
