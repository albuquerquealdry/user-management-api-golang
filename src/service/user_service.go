package service

import (
	"user-management/src/models"
	"user-management/src/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	// GetAllUsers() ([]models.User, error)
	// GetUserById(id uint) (*models.User, error)
	// UpdateUser(user *models.User) error
	// DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(user *models.User) error {
	// Lógica de validação ou processamento pode ser adicionada aqui
	// if user.Name == "" {
	// 	return errors.New("user name is required")
	// }
	// Chama o repositório para salvar no banco
	return s.userRepo.Create(user)
}
