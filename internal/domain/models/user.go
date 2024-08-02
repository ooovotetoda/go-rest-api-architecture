package models

type User struct {
	ID    int32  `json:"id" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Age   string `json:"age" validate:"required"`
}
