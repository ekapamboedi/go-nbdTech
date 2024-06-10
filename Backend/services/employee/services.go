package Employee

import (
	"github.com/ekapamboedi/go-nbdTech/services/employee/request"
)

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

	return nil
}
