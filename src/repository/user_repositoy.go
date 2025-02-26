package repository

import (
	"user-management/src/config"
	"user-management/src/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindById(id uint) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: config.DB,
	}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
