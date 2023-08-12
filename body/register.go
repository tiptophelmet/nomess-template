package body

// TODO: Need to follow "An Improved Handler" guidelines
// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

type Register struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=16"`
}
