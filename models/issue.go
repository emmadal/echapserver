package models

import (
	"echapserver/db"
)

// CreateIssue create isuue
func CreateIssue(issue Issues) error {
	query := "INSERT INTO issues (name, subject, user_name, phone, ticket_ref, user_id, description) VALUES (?,?,?,?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	result, err := stmt.Exec(issue.Name, issue.Subject, issue.UserName, issue.Phone, issue.TicketRef, issue.UserID, issue.Description)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	issue.ID = id
	return nil
}
