package Employee

import (
	"fmt"

	"github.com/ekapamboedi/go-nbdTech/model"
	"github.com/ekapamboedi/go-nbdTech/services/employee/request"
)

// var DB *gorm.DB

// func Query(identity helper.TokenPayload) (*[]model.User, error) {
// var resultData []model.User
// result := model.DB.Model(&model.User{}).Where("company_id = ?", identity.CompanyId).Find(&resultData)
// if result.Error != nil {
// 	return nil, result.Error
// }

// return &resultData, nil
// }

// func RegisterUser(req request.RequestCreateUser) error {
// var data model.User
// var count int64
// var jsonData []byte

// model.DB.Table("users").Limit(1)
// }

// func CreateOne(identity helper.TokenPayload, req request.RequestCreate) error {
func CreateOne(req request.RequestCreateEmployee) error {

	// define model
	data := model.Employee{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		// Created_at: req.Created_at,
		// Updated_at: req.Update_at,
	}
	result := model.DB.Table("Employee").Create(&data)

	if result.Error != nil {
		fmt.Print("result:", result)
		return result.Error
	}
	return nil
}
