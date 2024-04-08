package models

import (
	"echapserver/db"
)

// PremiumOffer allow user to subscribe/Unsubscribe to premium offer
func PremiumOffer(premium bool, userID int64) error {
	query := `UPDATE users SET premium = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query, )
	defer stmt.Close()

	if err != nil {
		return err
	}
	_, err = stmt.Exec(premium, userID)
	return err
}