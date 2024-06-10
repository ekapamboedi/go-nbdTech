package config

import (
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	supa "github.com/nedpals/supabase-go"
)

var (
	SERVER_PORT string
	ENVIRONMENT string

	SUPABASE_URL string
	SUPABASE_KEY string
	// DB_HOST     string
	// DB_PORT     string
	// DB_USER     string
	// DB_PASS     string
	// DB_NAME     string
)

func Init() {
	godotenv.Load(".env")

	SERVER_PORT = os.Getenv("SERVER_PORT")
	ENVIRONMENT = os.Getenv("ENVIRONMENT")

	SUPABASE_URL = os.Getenv("SUPABASE_URL")
	SUPABASE_KEY = os.Getenv("SUPABASE_KEY")

	// DB_PORT = os.Getenv("DB_PORT")
	// DB_NAME = os.Getenv("DB_NAME")
	// DB_USER = os.Getenv("DB_USER")
	// DB_PASS = os.Getenv("DB_PASS")
	// DB_HOST = os.Getenv("DB_HOST")

}

// func PgCon() (*sql.DB, error) {
// 	// Load konfigurasi dari file .env
// 	Init()

// 	// Format string koneksi ke PostgreSQL
// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 	DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

// 	// Buat koneksi ke database
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Uji koneksi ke database
// 	err = db.Ping()
// 	if err != nil {
// 		db.Close() // Tutup koneksi jika ping gagal
// 		return nil, err
// 	}

// 	fmt.Println("Berhasil terhubung ke database")

//		return db, nil
//	}
var Supabase *supa.Client

func SupaCon() {
	Init()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	supabase := supa.CreateClient(SUPABASE_URL, SUPABASE_KEY)
	if supabase != nil {
		log.Fatal("Error initializing Supabase client:", supabase)
	}

	// Jika tidak ada error, tampilkan pesan bahwa klien Supabase berhasil diinisialisasi
	fmt.Println("Supabase client initialized successfully")
}
