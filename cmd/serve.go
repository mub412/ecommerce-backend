package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()
	m := manager.With(middleware.Hudai, middleware.Logger)
	h := m(handlers.Test)
	mux.Handle("GET /uddin", h)
	mux.Handle("GET /mohi", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test))))

	mux.Handle("GET /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProducts))))
	mux.Handle("GET /products/{productId}", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProductByID))))
	mux.Handle("POST /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.CreateProducts))))

	globalRouter := global_router.GlobalRouter(mux)
	fmt.Println("Server running on : 8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
