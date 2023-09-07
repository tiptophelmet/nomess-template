package response

import "github.com/tiptophelmet/nomess-core/intl"

type registrationSuccessful struct {
	Message string
	Text    string
}

func RegistrationSuccessful() *registrationSuccessful {
	return &registrationSuccessful{
		Message: intl.Localize("registration_successful.message"),
		Text:    intl.Localize("registration_successful.text"),
	}
}
