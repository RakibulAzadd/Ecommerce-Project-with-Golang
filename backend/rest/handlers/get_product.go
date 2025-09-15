package handlers

import (
	"ecommerce/database/product_query"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid product ID", 400)
	}

	product, err := product_query.GetByIDFormDb(pId)

	if err != nil {
		util.SendError(w, "Error is found", 404)
		return
	}
	if product == nil {
		util.SendError(w, "Product not found", 404)
		return
	}

	util.SendData(w, product, 200)

}
