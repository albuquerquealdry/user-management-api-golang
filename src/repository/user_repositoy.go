package repository

import (
	"fmt"
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
	fmt.Println("teste Repository")
	return r.db.Create(user).Error
}

func (r *userRepository) FindById(id uint) (*models.User, error) {

	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
