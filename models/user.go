package models

import (
	"database/sql"
	"echapserver/db"
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
	query := "SELECT * FROM users WHERE phone = ? AND is_active = 1"
	row := db.DB.QueryRow(query, u.Phone)

	var biography, whatsapp, photo, tiktok, instagram sql.NullString

	var user User
	err := row.Scan(&user.ID, &user.Name, &biography, &user.Premium, &user.Phone, &user.IsActive, &user.Role, &user.CountryID, &user.CityID, &photo, &whatsapp, &tiktok, &instagram, &user.CreatedAt)

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
	var biography, photo, whatsapp, tiktok, instagram sql.NullString
	var user User

	if err := row.Scan(&user.ID, &user.Name, &biography, &user.Premium, &user.Phone, &user.IsActive, &user.Role, &user.CountryID, &user.CityID, &photo, &whatsapp, &tiktok, &instagram, &user.CreatedAt); err != nil {
		return nil, err
	}
	if biography.Valid || photo.Valid || whatsapp.Valid || tiktok.Valid || instagram.Valid {
		user.Biography = biography.String
		user.Photo = photo.String
		user.Whatsapp = whatsapp.String
		user.TikTok = tiktok.String
		user.Instagram = instagram.String
	}
	return &user, nil
}

// UpdateUser update a user
func UpdateUser(user User) error {
	query := `UPDATE users SET name = ?, biography = ?, premium = ?, phone = ?, country_id = ?, city_id = ?, photo = ?, whatsapp = ?, tiktok = ?, instagram = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Name, user.Biography, user.Premium, user.Phone, user.CountryID, user.CityID, user.Photo, user.Whatsapp, user.TikTok, user.Instagram, user.ID)

	return err
}

// DeleteAccount delete user account
func DeleteAccount(userID int64) error {
	query := `UPDATE users SET is_active = 0 WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID)

	return err
}
