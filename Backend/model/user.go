package model

import (
	"database/sql"
)

type User struct {
	UserId   sql.NullInt64  `json:"id"`
	UserName sql.NullString `json:"username"`
	Email    sql.NullString `json:"email"`
	Role     sql.NullString `json:"role"`
	Password sql.NullString `json:"password"`
	Session  sql.NullString `json:"session"`
}

func (User) TableName() string {
	return "users"
}
