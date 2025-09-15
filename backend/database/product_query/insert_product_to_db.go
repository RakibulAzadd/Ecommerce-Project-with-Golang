package product_query

import (
	"ecommerce/database"
)

func InsertProductToDb(p database.Product) error {

	lastID, err := GetLastID()
	if err != nil {
		return err
	}

	p.ID = lastID + 1

	query := `
		INSERT INTO products (id, title, description, price, img_url) 
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = db.Exec(query, p.ID, p.Title, p.Description, p.Price, p.ImgUrl)

	return err
}

func GetLastID() (int, error) {
	var lastID int
	err := db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM products").Scan(&lastID)
	return lastID, err
}
