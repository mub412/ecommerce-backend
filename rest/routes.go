package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /uddin",
		manager.With(
			http.HandlerFunc(handlers.Test),
			middlewares.Arekta,
		),
	)
	mux.Handle(
		"GET /mohi",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)
	mux.Handle(
		"GET /products/{productId}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)
}
