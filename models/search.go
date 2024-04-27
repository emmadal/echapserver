package models

import (
	"database/sql"
	"echapserver/db"
	"encoding/json"
)

// SearchArticleByName search all article by name and category_id
func SearchArticleByName(title string, id int64) ([]Article, error) {
	query := `SELECT * FROM article WHERE MATCH(title) AGAINST(+?) AND category_id = ? AND is_active = 1 ORDER BY created_at DESC`

	rows, err := db.DB.Query(query, title, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var articles []Article
	var photos sql.RawBytes
	for rows.Next() {
		var item Article
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.IsActive, &item.Phone, &item.Banner, &photos, &item.AuthorID, &item.CategoryID, &item.CountryID, &item.CityID, &item.CreatedAt)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(photos, &item.Photos)
		if err != nil {
			return nil, err
		}
		articles = append(articles, item)
	}
	return articles, nil
}
