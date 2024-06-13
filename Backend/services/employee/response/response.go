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
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Data    []ResponseCreateEmployee `json:"data"`
}

// Data
type ResponseCreateEmployee struct {
	// Id      int64  `json:"Id"`
	Name    string `jsom:"Name"`
	Phone   string `jsom:"Phone"`
	Email   string `jsom:"Email"`
	Address string `jsom:"Address"`
}

// Set

func (res *ResponseEmployee) Set(data []model.Employee) {
	for _, item := range data {
		res.Data = append(res.Data, ResponseCreateEmployee{
			// Id:      item.Id.Int64,
			Name:    item.Name,
			Phone:   item.Phone,
			Email:   item.Email,
			Address: item.Address,
		})
	}
}
