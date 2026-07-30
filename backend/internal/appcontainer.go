package app

import (
	"backend/internal/repository"
	"backend/internal/services"
)

type Container struct {
	Auth  *services.AuthService
	Crop  *services.CropService
	Order *services.OrderService
	AI    *services.AIService

	// Exposed so route-level auth middleware can look up the logged-in user.
	FarmerRepo *repository.FarmerRepository
}

func NewContainer(
	farmerRepo *repository.FarmerRepository,
	cropRepo *repository.CropRepository,
	orderRepo *repository.OrderRepository,
	diagnosisRepo *repository.DiagnosisRepository,
) *Container {
	return &Container{
		Auth:       services.NewAuthService(farmerRepo),
		Crop:       services.NewCropService(cropRepo),
		Order:      services.NewOrderService(orderRepo, cropRepo),
		AI:         services.NewAIService(diagnosisRepo),
		FarmerRepo: farmerRepo,
	}
}
