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
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	host := os.Getenv("HOST")

	credentials := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?parseTime=true", username, password, host, dbName)

	db, err := sql.Open("mysql", credentials)
	DB = db

	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	createTales()
}

func createTales() {
	start := time.Now()

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
			biography VARCHAR(200),
			premium TINYINT DEFAULT 0 NOT NULL,
			phone VARCHAR(20) NOT NULL,
			role TINYINT DEFAULT 0 NOT NULL,
			is_active TINYINT DEFAULT 1 NOT NULL,
			country_id INTEGER NOT NULL,
			city_id INTEGER NOT NULL,
			photo TEXT,
			whatsapp VARCHAR(20),
			tiktok VARCHAR(150),
			instagram VARCHAR(150),
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
			is_active TINYINT DEFAULT 1 NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(title)
		)`,

		`CREATE TABLE IF NOT EXISTS billing (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			label VARCHAR(30) NOT NULL,
			reference VARCHAR(30) NOT NULL,
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
			delivery INTEGER NOT NULL,
			is_active TINYINT DEFAULT 1 NOT NULL,
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

		`CREATE TABLE IF NOT EXISTS issues (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(100) NOT NULL,
			subject VARCHAR(100) NOT NULL,
			user_name VARCHAR(100) NOT NULL,
			phone VARCHAR(20) NOT NULL,
			ticket_ref VARCHAR(20) NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			description TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS chats (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			message TEXT NOT NULL,
			sender_id INTEGER NOT NULL,
			FOREIGN KEY (sender_id) REFERENCES users(id)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			receiver_id INTEGER NOT NULL,
			FOREIGN KEY (receiver_id) REFERENCES users(id)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	}
	for _, value := range queries {
		if _, err := DB.Exec(value); err != nil {
			panic(err)
		}
	}
	fmt.Println("Executed in:", time.Since(start))
}
