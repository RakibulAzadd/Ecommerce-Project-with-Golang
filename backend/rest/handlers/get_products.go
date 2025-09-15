package handlers

import (
	"ecommerce/database"
	"ecommerce/database/product_query"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	var products []database.Product

	products, err := product_query.GetAllListFormDb()
	if err != nil {
		util.SendData(w, "Their is some Error", 404)
	}

	util.SendData(w, products, 200)
}
