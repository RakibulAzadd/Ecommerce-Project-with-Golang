package product_query

import "ecommerce/database"

func GetAllListFormDb() ([]database.Product, error) {
	rows, err := db.Query("SELECT id, title, description, price, img_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []database.Product

	for rows.Next() {
		var p database.Product
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImgUrl)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
