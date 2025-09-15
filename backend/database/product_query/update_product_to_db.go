package product_query

import "ecommerce/database"

func UpdateProductFormDb(p database.Product) error {
	query := `
        UPDATE products 
        SET title=$1, description=$2, price=$3, img_url=$4 
        WHERE id=$5
    `
	_, err := db.Exec(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID)
	return err
}
