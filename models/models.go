package models

import "time"

// Article struct
type Article struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	Phone       string    `json:"phone" binding:"required"`
	Banner      string    `json:"banner" binding:"required"`
	Photos      []string  `json:"photos" binding:"required"`
	AuthorID    int64     `json:"author_id" binding:"required"`
	CategoryID  int64     `json:"category_id" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}

// User struct
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Biography string    `json:"biography" binding:"required"`
	Premium   bool      `json:"premium" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	Whatsapp  string    `json:"whatsapp"`
	TikTok    string    `json:"tiktok"`
	Instagram string    `json:"instagram"`
	CreatedAt time.Time `json:"createdAt"`
}

// Category struct
type Category struct {
	ID        int64  `json:"id"`
	Title     string `json:"title" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

// Config struct
type Config struct {
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD,unset"`
	Port     string `env:"PORT" envDefault:"3306"`
	Host     string `env:"HOST"`
	DBName   string `env:"DBNAME"`
}
