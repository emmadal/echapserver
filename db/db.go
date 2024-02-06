package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/joho/godotenv"
)

// DB database pointer
var DB *sql.DB

// InitDB initialize database
func InitDB() {
	defer recoverDB()

	credentials := parseEnv()
	db, err := sql.Open("mysql", credentials)
	DB = db

	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	go createTales()
}

func parseEnv() (credentials string) {
	defer recoverEnv()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	host := os.Getenv("HOST")
	credentials = fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?parseTime=true", username, password, host, dbName)
	return
}

func createTales() {
	defer recoverTable()
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(60) NOT NULL,
			biography TEXT NOT NULL,
			premium TINYINT DEFAULT 0 NOT NULL,
			phone VARCHAR(10) NOT NULL,
			whatsapp TEXT,
			tiktok TEXT,
			instagram TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	
		`CREATE TABLE IF NOT EXISTS category (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR(50) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS article (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR(60) NOT NULL,
			description TEXT NOT NULL,
			price INTEGER NOT NULL,
			phone VARCHAR(10) NOT NULL,
			banner TEXT NOT NULL,
			photos JSON,
			author_id  INTEGER NOT NULL,
			FOREIGN KEY (author_id) REFERENCES users(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			category_id INTEGER NOT NULL,
			FOREIGN KEY (category_id) REFERENCES category(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	}
	for _, value := range queries {
		_, err := DB.Exec(value)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Database creation in:", time.Since(time.Now()))
}
