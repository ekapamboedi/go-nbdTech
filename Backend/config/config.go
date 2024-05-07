package config

import (
	"database/sql"
	"fmt"
	"os"

	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

var (
	SERVER_PORT string
	ENVIRONMENT string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_NAME     string
)

// var OSS_ACCESS_KEY_SECRET string
// var OSS_ACCESS_KEY_ID string
// var OSS_BUCKET_NAME string
// var RABBITMQ_USERNAME string
// var RABBITMQ_PASSWORD string
// var RABBITMQ_HOST string
// var RABBITMQ_PORT string
// var OSS_REGION string
// var JWT_KEY string
// var JWT_ISSUER string

var ALLOWED_EXT = []string{"jpeg", "jpg", "png", "pdf", "docx", "xlsx"}

// var Bukcet *oss.Bucket
// var AMQPConn *amqp.Connection

func Init() {
	godotenv.Load(".env")

	SERVER_PORT = os.Getenv("SERVER_PORT")
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	// OSS_ACCESS_KEY_SECRET = os.Getenv("OSS_ACCESS_KEY_SECRET")
	// OSS_ACCESS_KEY_ID = os.Getenv("OSS_ACCESS_KEY_ID")
	// OSS_BUCKET_NAME = os.Getenv("OSS_BUCKET_NAME")
	// OSS_REGION = os.Getenv("OSS_REGION")
	// RABBITMQ_USERNAME = os.Getenv("RABBITMQ_USERNAME")
	// RABBITMQ_PASSWORD = os.Getenv("RABBITMQ_PASSWORD")
	// RABBITMQ_HOST = os.Getenv("RABBITMQ_HOST")
	// RABBITMQ_PORT = os.Getenv("RABBITMQ_PORT")
	// JWT_KEY = os.Getenv("JWT_KEY")
	// JWT_ISSUER = os.Getenv("DB_HOST")
}
func DbCon() (*sql.DB, error) {
	// Load konfigurasi dari file .env
	Init()

	// Format string koneksi ke PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	// Buat koneksi ke database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Uji koneksi ke database
	err = db.Ping()
	if err != nil {
		db.Close() // Tutup koneksi jika ping gagal
		return nil, err
	}

	fmt.Println("Berhasil terhubung ke database PostgreSQL")

	return db, nil
}

// func InitOss() {
// 	/// Obtain access credentials from environment variables. Before you run the sample code, make sure that the OSS_ACCESS_KEY_ID and OSS_ACCESS_KEY_SECRET environment variables are configured.
// 	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(-1)
// 	}

// 	// Create an OSSClient instance.
// 	// Specify the endpoint of the region in which the bucket is located. For example, if the bucket is located in the China (Hangzhou) region, set the endpoint to https://oss-cn-hangzhou.aliyuncs.com. Specify your actual endpoint.
// 	client, err := oss.New(OSS_REGION, "", "", oss.SetCredentialsProvider(&provider))
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(-1)
// 	}

// 	// Specify the name of the bucket. Example: examplebucket.
// 	bucket, err := client.Bucket(OSS_BUCKET_NAME)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(-1)
// 	}

// 	Bukcet = bucket
// }

// func InitAMQP() {
// 	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", RABBITMQ_USERNAME, RABBITMQ_PASSWORD, RABBITMQ_HOST, RABBITMQ_PORT))
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		conn.Close()
// 		return
// 	}

// 	fmt.Println("RabbitMQ Connection estabilished")

// 	AMQPConn = conn
// }
