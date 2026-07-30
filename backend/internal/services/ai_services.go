package services

import (
	"errors"
	"hash/fnv"

	"backend/internal/models"
	"backend/internal/repository"
)

type AIService struct {
	repo *repository.DiagnosisRepository
}

func NewAIService(repo *repository.DiagnosisRepository) *AIService {
	return &AIService{repo: repo}
}

// possibleFindings is a placeholder result set. There's no real image model
// wired up yet — swap the body of Diagnose for a call to a real vision
// model (e.g. a crop-disease API, or Claude's vision) when you're ready.
// Nothing else in the app needs to change; the handler, template, and
// dashboard count all just read whatever Diagnose returns.
var possibleFindings = []string{
	"Leaves look healthy — no signs of disease detected.",
	"Early signs of leaf blight detected. Consider a copper-based fungicide.",
	"Possible nitrogen deficiency — leaves show yellowing from the base.",
	"Signs of pest damage, likely armyworm. Inspect the underside of leaves.",
	"Mild fungal spotting detected. Improve airflow and avoid overwatering.",
}

func (s *AIService) Diagnose(farmerID int, imageName string, imageBytes []byte) (*models.Diagnosis, error) {
	if len(imageBytes) == 0 {
		return nil, errors.New("no image was uploaded")
	}

	h := fnv.New32a()
	h.Write(imageBytes)
	result := possibleFindings[int(h.Sum32())%len(possibleFindings)]

	diagnosis := &models.Diagnosis{
		FarmerID:  farmerID,
		ImageName: imageName,
		Result:    result,
	}

	if err := s.repo.Create(diagnosis); err != nil {
		return nil, err
	}

	return diagnosis, nil
}

func (s *AIService) History(farmerID int) ([]models.Diagnosis, error) {
	return s.repo.ListByFarmer(farmerID)
}

func (s *AIService) CountThisMonth(farmerID int) (int, error) {
	return s.repo.CountThisMonth(farmerID)
}
