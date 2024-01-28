package model

type Twoot struct {
	Id       uint
	Contents string `form:"contents" validate:"required,min=1,max=1000"`
	UserId   uint
}
