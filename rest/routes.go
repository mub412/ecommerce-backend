package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	)
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
			middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProducts),
			middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProducts),
			middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
			middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)
}
