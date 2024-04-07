package models

import "echapserver/db"

// FindCitiesByCountryID get all cities by country
func FindCitiesByCountryID(countryID int64) ([]City, error) {
	query := `SELECT * FROM city WHERE country_id = ? ORDER BY label ASC`
	rows, err := db.DB.Query(query, countryID)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	var cities []City

	for rows.Next() {
		var city City
		err := rows.Scan(&city.ID, &city.Label, &city.CountryID, &city.CreatedAt)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, nil
}
