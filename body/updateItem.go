package body

type UpdateItem struct {
	Name        string `validate:"required"`
	Description string `validate:"required,min=50"`
}
