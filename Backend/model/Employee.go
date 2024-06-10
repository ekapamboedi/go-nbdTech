package model

type Employee struct {
	Id         int64  `json:"Id"`
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Phone      string `json:"Phone"`
	Address    string `json:"Address"`
	Created_at string `json:"Created_at"`
	Updated_at string `json:"Update_at"`
}

func (Employee) TableName() string {
	return "Employee"
}
