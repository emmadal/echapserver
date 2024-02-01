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

// GetCategoryByID get a category
func GetCategoryByID(categoryID int64) (*Category, error) {
	query := `SELECT * FROM category WHERE ID = ?`
	rows := db.DB.QueryRow(query, categoryID)
	var category Category

	err := rows.Scan(&category.ID, &category.Title, &category.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// UpdateCategory update a category
func UpdateCategory(category Category) error {
	query := `UPDATE category SET title = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()

	_, err = stmt.Exec(category.Title, category.ID)
	return err
}
