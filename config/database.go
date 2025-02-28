package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah instance database global
var DB *gorm.DB

// ConnectDB menghubungkan aplikasi ke database PostgreSQL
func ConnectDB() {
	// Load variabel dari .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Konfigurasi DSN PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Inisialisasi koneksi database PostgreSQL dan simpan ke variabel global DB
	var errDB error
	DB, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("Failed to connect to database:", errDB)
	}

	fmt.Println("Database connected successfully!")
}
