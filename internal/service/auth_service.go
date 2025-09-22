package service

import (
	"errors"
	"umkmgubernur/internal/dto"
	"umkmgubernur/internal/models"
	"umkmgubernur/internal/repository"
)

type AuthService interface {
	RegisterUser(req dto.RegisterRequestDTO) (*models.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) RegisterUser(req dto.RegisterRequestDTO) (*models.User, error) {
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}
	// Hash password, create user, save to DB
	newUser := &models.User{
		FullName: req.FullName,
		Email: req.Email,
		PasswordHash: req.Password, // Replace with hash
	}
	if err := s.userRepo.Create(newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}
