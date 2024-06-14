package model

import (
	"fmt"
	"log"
	"time"

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

	database, err := gorm.Open(postgres.Open(connStrSupa), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("failed to get database connection pool: %v", err)
	}

	// Set maximum number of open connections to the database
	sqlDB.SetMaxOpenConns(25)

	// Set maximum number of idle connections to the database
	sqlDB.SetMaxIdleConns(25)

	// Set the maximum lifetime of a connection
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DB = database
}
