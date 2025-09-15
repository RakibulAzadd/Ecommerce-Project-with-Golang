package product_query

import "fmt"

func DeleteProductFormDb(productID int) error {
	query := "DELETE FROM products WHERE id=$1"

	res, err := db.Exec(query, productID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		// কোনো row delete হয়নি মানে ঐ ID পাওয়া যায়নি
		return fmt.Errorf("no product found with id %d", productID)
	}

	return nil
}
