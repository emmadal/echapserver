package models

import "oblackserver/db"

// FindCountries get all countries
func FindCountries() ([]Country, error) {
	query := `SELECT * FROM country ORDER BY label ASC`
	rows, err := db.DB.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	var countries []Country

	for rows.Next() {
		var country Country
		err := rows.Scan(&country.ID, &country.Label, &country.CreatedAt)
		if err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}
	return countries, nil
}
