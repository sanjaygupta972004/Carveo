package services

import (
	"carveo/models"
	"carveo/repositories"
)

type UserService interface {
	RegisterUser(user models.User) (models.User, error)
	LoginUser(email, password string) (models.User, error)
	GetUserProfile(id string) (models.User, error)
	UpdateUserProfile(id string, user models.User) (models.User, error)
	DeleteUserProfile(id string) error
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) LoginUser(email string, password string) (models.User, error) {
	panic("unimplemented")
}

func (u *userService) RegisterUser(user models.User) (models.User, error) {
	panic("unimplemented")
}

func (u *userService) GetAllUsers() ([]models.User, error) {
	panic("unimplemented")
}

func (u *userService) GetUserProfile(id string) (models.User, error) {
	panic("unimplemented")
}

func (u *userService) UpdateUserProfile(id string, user models.User) (models.User, error) {
	panic("unimplemented")
}

func (u *userService) DeleteUserProfile(id string) error {
	panic("unimplemented")
}
