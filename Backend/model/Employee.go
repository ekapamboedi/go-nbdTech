package model

type Employee struct {
	// Id         sql.NullInt64  `json:"Id"`
	Name    string `json:"Name" gorm:"column:Name"`
	Phone   string `json:"Phone" gorm:"column:Phone"`
	Email   string `json:"Email" gorm:"column:Email"`
	Address string `json:"Address" gorm:"column:Address"`

	// Created_at sql.NullString `json:"Created_at"`
	// Updated_at sql.NullString `json:"Update_at"`
}

func (Employee) TableName() string {
	return "Employee"
}
