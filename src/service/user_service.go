package service

import (
	"fmt"
	"strconv"
	"sync"
	"user-management/src/models"
	"user-management/src/repository"
	"user-management/src/utils"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
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
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	var hashedPassword string
	var cpfIsValid bool

	wg.Add(2)

	go utils.HashPassword(user.Password, &wg, &hashedPassword, errChan)

	go utils.IsValidCPF(strconv.Itoa(user.CPF), &wg, &cpfIsValid, errChan)

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}
	if !cpfIsValid {
		return fmt.Errorf("CPF inválido")
	}

	user.Password = hashedPassword
	return s.userRepo.Create(user)
}

func (s *userService) GetUserById(id uint) (*models.User, error) {
	return s.userRepo.FindById(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
