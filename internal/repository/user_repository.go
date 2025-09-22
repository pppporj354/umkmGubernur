package repository

import (
	"umkmgubernur/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db any // Replace 'any' with *gorm.DB when GORM is imported
}

func NewUserRepository(db any) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	// Implement DB logic here
	return nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	// Implement DB logic here
	return nil, nil
}
