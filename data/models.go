// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package data

import (
	"database/sql"
)

type User struct {
	ID          int64
	Username    string
	Email       string
	Password    string
	Bio         sql.NullString
	DisplayName sql.NullString
}