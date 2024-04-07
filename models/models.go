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
	CountryID   int64     `json:"country_id" binding:"required"`
	CityID      int64     `json:"city_id" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

// User struct
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Biography string    `json:"biography"`
	Premium   bool      `json:"premium"`
	Phone     string    `json:"phone" binding:"required"`
	CountryID int64     `json:"country_id" binding:"required"`
	CityID    int64     `json:"city_id" binding:"required"`
	Photo     string    `json:"photo"`
	Whatsapp  string    `json:"whatsapp"`
	TikTok    string    `json:"tiktok"`
	Instagram string    `json:"instagram"`
	CreatedAt time.Time `json:"created_at"`
}

// Country struct
type Country struct {
	ID        int64     `json:"id"`
	Label     string    `json:"label" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// City struct
type City struct {
	ID        int64     `json:"id"`
	Label     string    `json:"label" binding:"required"`
	CountryID int64     `json:"country_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// Category struct
type Category struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" binding:"required"`
	UserID    int64     `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// Billing struct
type Billing struct {
	ID          int64     `json:"id"`
	Label       string    `json:"label" binding:"required"`
	Price       int64     `json:"price" binding:"required"`
	UserID      int64     `json:"user_id" binding:"required"`
	PaymentDate time.Time `json:"payment_date"`
	CreatedAt   time.Time `json:"created_at"`
}

// Config struct
type Config struct {
	Username    string `env:"USERNAME"`
	Password    string `env:"PASSWORD,unset"`
	Port        string `env:"PORT" envDefault:"3306"`
	Host        string `env:"HOST"`
	DBName      string `env:"DBNAME"`
	CloudName   string `env:"CloudName"`
	CloudSecret string `env:"CloudSecret"`
	CloudAPIKey string `env:"CloudAPIKey"`
}

// AuthStruct struct
type AuthStruct struct {
	Phone string `json:"phone" binding:"required"`
}

// OTP struct
type OTP struct {
	ID         int64     `json:"id"`
	Code       string    `json:"code" binding:"required"`
	IsUsed     bool      `json:"is_used"`
	Expiration time.Time `json:"expiration"`
	UserID     int64     `json:"user_id" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
}

// OTPAuth struct
type OTPAuth struct {
	Code string `json:"code"`
}
