package models

import (
	"database/sql"
	"oblackserver/db"
)

// CreateUser create user
func CreateUser(user User) error {
	query := "INSERT INTO users (name, phone, country_id, city_id) VALUES (?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Name, user.Phone, user.CountryID, user.CityID)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = userID
	return nil
}

// LoginUser sign in user
func LoginUser(u AuthStruct) (*User, error) {
	query := "SELECT * FROM users WHERE phone = ?"
	row := db.DB.QueryRow(query, u.Phone)

	var biography, whatsapp, photo, tiktok, instagram sql.NullString

	var user User
	err := row.Scan(&user.ID, &user.Name, &biography, &user.Premium, &user.Phone, &photo, &whatsapp, &tiktok, &instagram, &user.CreatedAt)

	if biography.Valid && whatsapp.Valid && tiktok.Valid && instagram.Valid {
		user.Biography = ""
		user.Photo = ""
		user.Whatsapp = ""
		user.TikTok = ""
		user.Instagram = ""
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUserByID get user data by his ID
func FindUserByID(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var biography, whatsapp, photo, tiktok, instagram sql.NullString

	var user User
	err := row.Scan(&user.ID, &user.Name, &biography, &user.Premium, &user.Phone, &user.CountryID, &user.CityID, &photo, &whatsapp, &tiktok, &instagram, &user.CreatedAt)

	if biography.Valid && whatsapp.Valid && tiktok.Valid && instagram.Valid {
		user.Biography = ""
		user.Photo = ""
		user.Whatsapp = ""
		user.TikTok = ""
		user.Instagram = ""
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
