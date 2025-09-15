package handlers

import (
	"ecommerce/database"
	"ecommerce/database/product_query"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON", 400)
	}

	err = product_query.InsertProductToDb(newProduct)

	if err != nil {
		util.SendData(w, "do not inserted product successfully", 201)
	}

	util.SendData(w, "Successfully Product inserted", 201)
}
