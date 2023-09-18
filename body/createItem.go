package body

// TODO: Need to follow "An Improved Handler" guidelines
// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

type CreateItem struct {
	Name        string `validate:"required"`
	Description string `validate:"required,min=50"`
}
