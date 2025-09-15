package product_query

import (
	"database/sql"
	"ecommerce/database"
)

func GetByIDFormDb(productID int) (*database.Product, error) {
	var p database.Product

	err := db.QueryRow(
		"SELECT id, title, description, price, img_url FROM products WHERE id=$1",
		productID,
	).Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImgUrl)

	if err == sql.ErrNoRows {
		// product পাওয়া যায়নি
		return nil, nil
	} else if err != nil {
		// অন্য কোনো error
		return nil, err
	}

	return &p, nil
}
