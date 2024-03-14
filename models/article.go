package models

import (
	"encoding/json"
	"oblackserver/db"
)

// CreateArticle create article
func CreateArticle(article Article) error {
	query := `INSERT INTO article(title, description, price, phone, banner, photos, author_id, category_id, country_id, city_id) VALUES (?,?,?,?,?,?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	photos, err := json.Marshal(article.Photos)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(article.Title, article.Description, article.Price, article.Phone, article.Banner, photos, article.AuthorID, article.CategoryID, article.CountryID, article.CityID)

	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	article.ID = id

	return err
}