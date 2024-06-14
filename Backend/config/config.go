package config

import (
	"database/sql"
	"fmt"
	"os"

	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	SERVER_PORT string
	ENVIRONMENT string

	SUPABASE_URL string
	SUPABASE_KEY string
	DB_NAME      string
	DB_USER      string
	DB_HOST      string
	DB_PORT      string
	DB_PASS      string
	DB           *sql.DB
)

func Init() {
	godotenv.Load(".env")

	SERVER_PORT = os.Getenv("SERVER_PORT")
	ENVIRONMENT = os.Getenv("ENVIRONMENT")

	SUPABASE_URL = os.Getenv("SUPABASE_URL")
	SUPABASE_KEY = os.Getenv("SUPABASE_KEY")

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
}

// koneksi pakai postgre
func PgCon() (*sql.DB, error) {
	// Load konfigurasi dari file .env
	Init()

	// Format string koneksi ke PostgreSQL
	connStrPG :=
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
			DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	fmt.Print("PGCon:", connStrPG)

	// Buat koneksi ke database
	db, err := sql.Open("postgres", connStrPG)
	if err != nil {
		return nil, err
	}

	// Uji koneksi ke database
	err = db.Ping()
	if err != nil {
		db.Close() // Tutup koneksi jika ping gagal
		return nil, err
	}

	fmt.Println("Berhasil terhubung ke database")

	return db, nil
}

// koneksi pake client supabase
// var Supabase *supa.Client
func SupaCon() (*sql.DB, error) {
	connStrSupa := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	db, err := sql.Open("postgres", connStrSupa)
	if err != nil {
		return nil, err
		db.Close()
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	db.Close()
	fmt.Println("Supabase client initialized successfully")

	// Querying the column names
	// query := "SELECT column_name FROM information_schema.columns WHERE table_name = 'Employee';"
	// rows, err := db.Query(query)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// fmt.Println("Columns in Employee table:")
	// for rows.Next() {
	// 	var columnName string
	// 	if err := rows.Scan(&columnName); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(columnName)
	// }
	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	return db, nil
}
