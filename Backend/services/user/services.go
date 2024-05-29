package User

import "github.com/ekapamboedi/go-project/services/user/request"

// import "github.com/ekapamboedi/go-project/services/user/request"

// "github.com/ekapamboedi/go-project/helper"
// "github.com/ekapamboedi/go-project/model"
// "github.com/ekapamboedi/go-project/database"
// "github.com/ekapamboedi/go-project/services/user/request"
// viewmodel "github.com/ekapamboedi/go-project/viewModel"

// func Query(identity helper.TokenPayload) (*[]model.User, error) {
// 	var resultData []model.User
// 	result := model.DB.Model(&model.User{}).Where("company_id = ?", identity.CompanyId).Find(&resultData)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &resultData, nil
// }

func RegisterUser(req request.RequestCreateUser) error {
	// var data model.User
	// var count int64
	// var jsonData []byte

	// model.DB.Table("users").Limit(1)
}
