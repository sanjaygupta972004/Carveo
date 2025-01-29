package controllers

import (
	"carveo/models"
	"carveo/services"
)

type UserController interface {
	RegisterUser(user models.User) (models.User, error)
	LoginUser(email, password string) (models.User, error)
	GetUserProfile(id string) (models.User, error)
	UpdateUserProfile(id string, user models.User) (models.User, error)
	DeleteUserProfile(id string) error
	GetAllUsers() ([]models.User, error)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (u *userController) GetAllUsers() ([]models.User, error) {
	panic("unimplemented")
}

func (u *userController) GetUserProfile(id string) (models.User, error) {
	panic("unimplemented")
}

func (u *userController) LoginUser(email string, password string) (models.User, error) {
	panic("unimplemented")
}

func (u *userController) RegisterUser(user models.User) (models.User, error) {
	panic("unimplemented")
}

func (u *userController) UpdateUserProfile(id string, user models.User) (models.User, error) {
	panic("unimplemented")
}

func (u *userController) DeleteUserProfile(id string) error {
	panic("unimplemented")
}
