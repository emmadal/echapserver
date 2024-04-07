package db

import (
	"database/sql"
	"fmt"
	"log"
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

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	host := os.Getenv("HOST")

	credentials := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?parseTime=true", username, password, host, dbName)

	db, err := sql.Open("mysql", credentials)
	DB = db

	if err != nil {
		log.Fatalln(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	createTales()
}

func createTales() {
	defer recoverTable()
	queries := []string{
		`CREATE TABLE IF NOT EXISTS country (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			label VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(label)
		)`,

		`CREATE TABLE IF NOT EXISTS city (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			label VARCHAR(100) NOT NULL,
			country_id INTEGER NOT NULL,
			FOREIGN KEY (country_id) REFERENCES country(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(label)
		)`,

		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(100) NOT NULL,
			biography VARCHAR(255),
			premium TINYINT DEFAULT 0 NOT NULL,
			phone VARCHAR(15) NOT NULL,
			country_id INTEGER NOT NULL,
			city_id INTEGER NOT NULL,
			photo TEXT,
			whatsapp VARCHAR(15),
			tiktok TEXT,
			instagram TEXT,
			FOREIGN KEY (city_id) REFERENCES city(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			FOREIGN KEY (country_id) REFERENCES country(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(phone)
		)`,

		`CREATE TABLE IF NOT EXISTS category (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR(100) NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(title)
		)`,

		`CREATE TABLE IF NOT EXISTS billing (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			label VARCHAR(255) NOT NULL,
			price INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS article (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			price INTEGER NOT NULL,
			phone VARCHAR(15) NOT NULL,
			banner TEXT NOT NULL,
			photos JSON NOT NULL,
			author_id  INTEGER NOT NULL,
			FOREIGN KEY (author_id) REFERENCES users(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			category_id INTEGER NOT NULL,
			FOREIGN KEY (category_id) REFERENCES category(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			country_id INTEGER NOT NULL,
			FOREIGN KEY (country_id) REFERENCES country(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			city_id INTEGER NOT NULL,
			FOREIGN KEY (city_id) REFERENCES city(id) 
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS otp (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			code VARCHAR(7) NOT NULL,
			is_used TINYINT DEFAULT 0 NOT NULL,
			expiration TIMESTAMP,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) 
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
