package model

type User struct {
	Id          uint
	Username    string `form:"username" validate:"required,min=2,max=50"`
	Email       string `form:"email" validate:"required,email"`
	Password    string `form:"password" validate:"required"`
	DisplayName string `form:"display_name" validate:"omitempty,min=1,max=50"`
	Bio         string `form:"bio" validate:"max=200"`
}
