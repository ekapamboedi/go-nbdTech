package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ekapamboedi/go-nbdTech/config"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() {
	connStrSupa := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	database, err := gorm.Open(postgres.Open(connStrSupa))
	if err != nil {
		fmt.Println(err.Error())
	}
	connPool, _ := database.DB()
	connPool.SetMaxOpenConns(5)

	DB = database
}
