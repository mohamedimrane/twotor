package model

import "time"

type Twoot struct {
	Id        uint
	Contents  string `form:"contents" validate:"required,min=1,max=1000"`
	CreatedAt time.Time
	UserId    uint
}
