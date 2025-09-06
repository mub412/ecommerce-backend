package main

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/product"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create", http.HandlerFunc(handlers.CreateProducts))
	globalRouter := global_router.GlobalRouter(mux)
	fmt.Println("Server running on : 8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := product.Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       100,
		ImgUrl:      "https://i.chaldn.com/_mpimage/komola-orange-imported-50-gm-1-kg?src=https%3A%2F%2Feggyolk.chaldal.com%2Fapi%2FPicture%2FRaw%3FpictureId%3D64292&q=best&v=1&m=1600",
	}
	prd2 := product.Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is Green",
		Price:       150,
		ImgUrl:      "https://assets.clevelandclinic.org/transform/LargeFeatureImage/cd71f4bd-81d4-45d8-a450-74df78e4477a/Apples-184940975-770x533-1_jpg",
	}
	prd3 := product.Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Yellow",
		Price:       50,
		ImgUrl:      "https://www.orchardfood.co.za/cdn/shop/files/Bananas-2.png?v=1753697464&width=1946",
	}
	prd4 := product.Product{
		ID:          4,
		Title:       "Lichu",
		Description: "Lichu is red",
		Price:       10,
		ImgUrl:      "https://media.gettyimages.com/id/566454679/photo/lychee-fruits.jpg?s=1024x1024&w=gi&k=20&c=mb13OLKoGK9QZYHi3j5nMgBfkuhSMue19wMRCrLWy5g=",
	}

	product.ProductList = append(product.ProductList, prd1)
	product.ProductList = append(product.ProductList, prd2)
	product.ProductList = append(product.ProductList, prd3)
	product.ProductList = append(product.ProductList, prd4)

}
