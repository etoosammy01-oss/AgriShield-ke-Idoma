package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.FarmerRepository
}

func NewAuthService(repo *repository.FarmerRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(firstName, lastName, phone, password, location string) error {

	if firstName == "" || lastName == "" {
		return errors.New("name is required")
	}

	if phone == "" {
		return errors.New("phone is required")
	}

	if password == "" {
		return errors.New("password is required")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	existing, err := s.repo.GetByPhone(phone)
	if err != nil {
		return err
	}

	if existing != nil {
		return errors.New("phone number already registered")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	farmer := &models.Farmer{
		FullName: firstName + " " + lastName,
		Phone: phone,
		PasswordHash: string(hash),
		Location: location,
	}

	return s.repo.Create(farmer)
}