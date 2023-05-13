package dto

type CreateBookRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}
