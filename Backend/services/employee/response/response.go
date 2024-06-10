package response

import (
	"github.com/ekapamboedi/go-nbdTech/model"
)

// Core Response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Response List
type ResponseEmployee struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []GetEmployee `json:"data"`
}

// Data
type GetEmployee struct {
	Id      int64  `json:"Id"`
	Name    string `jsom:"Name"`
	Email   string `jsom:"Email"`
	Address string `jsom:"Address"`
	Phone   string `jsom:"Phone"`
}

// Set

func (res *ResponseEmployee) Set(data []model.Employee) {
	for _, item := range data {
		res.Data = append(res.Data, GetEmployee{
			Id:      item.Id,
			Name:    item.Name,
			Email:   item.Email,
			Address: item.Address,
			Phone:   item.Phone,
		})
	}
}
