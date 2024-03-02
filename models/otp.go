package models

import "oblackserver/db"

// GetOTPCodeByUserID get a otp code
func GetOTPCodeByUserID(userID int64) (*OTP, error) {
	query := `SELECT id, code, is_used, expiration, user_id FROM otp WHERE user_id = ?`
	rows := db.DB.QueryRow(query, userID)

	var otp OTP

	err := rows.Scan(&otp.ID, &otp.Code, &otp.IsUsed, &otp.Expiration, &otp.UserID)
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

// SaveOTPCode save the otp code in database
func SaveOTPCode(otp OTP) error {
	query := `INSERT INTO otp(code, expiration, user_id) VALUES(?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	result, err := stmt.Exec(otp.Code, otp.Expiration, otp.UserID)
	id, err := result.LastInsertId()
	otp.ID = id

	return err
}

// UpdateOTPCode update a otp code
func UpdateOTPCode(otp OTP) error {
	query := `UPDATE otp SET is_used = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()

	_, err = stmt.Exec(otp.IsUsed, otp.ID)
	return err
}
