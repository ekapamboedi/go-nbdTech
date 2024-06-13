package model

import (
	"database/sql"
)

type ErrorLog struct {
	Id           int64          `gorm:"primaryKey" json:"id"`
	Endpoint     sql.NullString `json:"endpoint"`
	Method       sql.NullString `json:"method"`
	Username     sql.NullString `json:"username"`
	Role         sql.NullString `json:"role"`
	ErrorMessage sql.NullString `json:"errorMessage"`
	Description  sql.NullString `json:"description"`
	CreatedAt    sql.NullTime   `json:"createdAt"`
}

func (ErrorLog) TableName() string {
	return "error_logs"
}
