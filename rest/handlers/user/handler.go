package user

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type Handler struct {
	//middlewares *middlewares.Manager
	cnf      *config.Config
	userRepo repo.UserRepo
}

func NewHandler(cnf *config.Config, userRepo repo.UserRepo) *Handler {
	return &Handler{
		cnf:      cnf,
		userRepo: userRepo,
	}
}
