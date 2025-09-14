package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middlewares.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler(middlewares)
	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()

}
