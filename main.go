package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadMySqlDBCreds() {
	// Initialize .env vars
	dbName := os.Getenv("MYSQL_DBNAME")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")

	fmt.Println("dbName: ", dbName)
	fmt.Println("dbUser: ", dbUser)
	fmt.Println("dbPassword: ", dbPassword)
	fmt.Println("dbHost: ", dbHost)
	fmt.Println("dbPort: ", dbPort)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("dsn: ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection failed!")
	}

	DB = db
	fmt.Println("DB connection successful!")
}

func main() {
	// Load .env
	godotenv.Load()

	// Load MySQL DB details
	loadMySqlDBCreds()
}
