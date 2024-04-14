package models

import "echapserver/db"

// GetAllCategories return the list of categories
func GetAllCategories() ([]Category, error) {
	rows, err := db.DB.Query("SELECT id, title FROM category WHERE is_active = 1 ORDER BY title")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var categories []Category
	for rows.Next() {
		var cat Category
		err := rows.Scan(&cat.ID, &cat.Title)
		if err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

// CreateCategory create category
func CreateCategory(category Category) error {
	query := `INSERT INTO category(title, user_id) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	result, err := stmt.Exec(category.Title, category.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	category.ID = id

	return err
}

// GetCategoryByID get a category
func GetCategoryByID(categoryID int64) (*Category, error) {
	query := `SELECT id, title, user_id FROM category WHERE ID = ? AND is_active = 1`
	rows := db.DB.QueryRow(query, categoryID)
	var category Category

	err := rows.Scan(&category.ID, &category.Title, &category.UserID)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// UpdateCategory update a category
func UpdateCategory(category Category) error {
	query := `UPDATE category SET title = ? WHERE id = ? AND is_active = 1`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(category.Title, category.ID)
	return err
}

// DeleteCategory delete a category
func DeleteCategory(categoryID int64) error {
	query := `UPDATE category set is_active = 0 WHERE ID = ?`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(categoryID)
	return err
}
