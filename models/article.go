package models

import (
	"database/sql"
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

// GetAllArticle fetch article by categoryId
func GetAllArticle(articleID string) ([]Article, error) {
	query := `SELECT * FROM ARTICLE WHERE category_id = ? ORDER BY created_at DESC`

	rows, err := db.DB.Query(query, articleID)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var articles []Article
	var photos sql.RawBytes

	for rows.Next() {
		var item Article
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.Phone, &item.Banner, &photos, &item.AuthorID, &item.CategoryID, &item.CountryID, &item.CityID, &item.CreatedAt)
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

// FindArticleByID get a article by articleID
func FindArticleByID(articleID int64) (*Article, error) {
	query := `SELECT * FROM ARTICLE WHERE ID = ?`
	row := db.DB.QueryRow(query, articleID)

	var item Article
	var photos []byte

	err := row.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.Phone, &item.Banner, &photos, &item.AuthorID, &item.CategoryID, &item.CountryID, &item.CityID, &item.CreatedAt)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(photos, &item.Photos)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// DeleteArticle delete an article
func DeleteArticle(articleID int64) error {
	query := `DELETE FROM article WHERE ID = ?`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(articleID)
	return err
}

// GetArticlesByUser fetch all articles by userID
func GetArticlesByUser(userID, offset int64) ([]Article, error) {
	query := `SELECT * FROM ARTICLE WHERE author_id = ? ORDER BY created_at DESC LIMIT 10 OFFSET ?`

	rows, err := db.DB.Query(query, userID, offset)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var articles []Article
	var photos sql.RawBytes

	for rows.Next() {
		var item Article
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.Phone, &item.Banner, &photos, &item.AuthorID, &item.CategoryID, &item.CountryID, &item.CityID, &item.CreatedAt)
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
