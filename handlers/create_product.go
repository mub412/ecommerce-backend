package handlers

import (
	"ecommerce/product"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {

	var newProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", 400)
		return
	}
	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)
	util.SendData(w, newProduct, 201)
}
