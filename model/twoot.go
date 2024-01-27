package model

type Twoot struct {
	Id       uint
	Contents string `form:"contents" validate:"required,min=2,max=50"`
	UserId   uint
}
