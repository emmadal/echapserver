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
		var category Category
		err := rows.Scan(&category.ID, &category.Title, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
