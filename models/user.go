package models

import "oblackserver/db"

// CreateUser create user
func CreateUser(user User) error {
	query := "INSERT INTO users (name, phone) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Name, user.Phone)
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
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Biography, &user.Premium, &user.Phone, &user.Whatsapp, &user.TikTok, &user.Instagram, &user.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil
}