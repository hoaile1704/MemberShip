package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

// InitDB khởi tạo kết nối cơ sở dữ liệu linh hoạt dựa trên loại cơ sở dữ liệu
func InitDB() *gorm.DB {
	// Load biến môi trường từ file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Không thể tải file .env, sử dụng biến môi trường có sẵn")
	}

	// Lấy loại cơ sở dữ liệu từ biến môi trường
	dbType := os.Getenv("DB_TYPE")

	var db *gorm.DB

	switch dbType {
	case "postgres":
		db = connectPostgres()
	case "mysql":
		db = connectMySQL()
	case "sqlite":
		db = connectSQLite()
	default:
		log.Fatalf("Loại cơ sở dữ liệu không được hỗ trợ: %s", dbType)
	}

	return db
}

// Kết nối đến PostgreSQL
func connectPostgres() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Chuỗi kết nối PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Lỗi kết nối PostgreSQL: %v", err)
	}

	return db
}

// Kết nối đến MySQL
func connectMySQL() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Chuỗi kết nối MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Lỗi kết nối MySQL: %v", err)
	}

	return db
}

// Kết nối đến SQLite
func connectSQLite() *gorm.DB {
	sqlitePath := os.Getenv("DB_SQLITE_PATH")

	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Lỗi kết nối SQLite: %v", err)
	}

	return db
}
