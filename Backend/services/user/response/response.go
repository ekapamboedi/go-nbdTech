package response

import (
	"github.com/ekapamboedi/go-project/model"
)

// Core Response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Response List
type ResponseUser struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []GetUser `json:"data"`
}

// Data
type GetUser struct {
	UserId   int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// Set

func (res *ResponseUser) Set(data []model.User) {
	for _, item := range data {
		res.Data = append(res.Data, GetUser{
			UserId:   item.UserId.Int64,
			UserName: item.UserName.String,
			Email:    item.Email.String,
			Role:     item.Role.String,
		})
	}
}
