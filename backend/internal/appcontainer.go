package app

import (
	"backend/internal/repository"
	"backend/internal/services"
)

type Container struct {
	Auth *services.AuthService
}

func NewContainer(repo *repository.FarmerRepository) *Container {

	return &Container{
		Auth: services.NewAuthService(repo),
	}

}
