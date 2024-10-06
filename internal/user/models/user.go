package models

type User struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name" validate:"required"`
	Age   int32  `json:"age" validate:"gt=0"`
}
