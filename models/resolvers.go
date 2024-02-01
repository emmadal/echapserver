package models

import "oblackserver/db"

// GetAllCategories return the list of categories
func GetAllCategories() ([]Category, error) {
	rows, err := db.DB.Query("SELECT * FROM category")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var categories []Category
	for rows.Next() {
		var cat Category
		err := rows.Scan(&cat.ID, &cat.Title, &cat.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

// CreateCategory create category
func CreateCategory(category Category) error {
	query := `INSERT INTO category(title) VALUES (?)`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	result, err := stmt.Exec(category.Title)
	id, err := result.LastInsertId()
	category.ID = id

	return err
}
