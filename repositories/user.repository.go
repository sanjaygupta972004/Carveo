package repositories

import (
	"carveo/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	LoginUser(email, password string) (models.User, error)
	GetUserByID(id uuid.UUID) (models.User, error)
	UpdateUser(id uuid.UUID, user models.User) (models.User, error)
	DeleteUser(id uuid.UUID) error
	GetAllUsers() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(user models.User) (models.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetAllUsers() ([]models.User, error) {
	panic("unimplemented")
}

func (u *userRepository) GetUserByID(id uuid.UUID) (models.User, error) {
	panic("unimplemented")
}

func (u *userRepository) LoginUser(email string, password string) (models.User, error) {
	panic("unimplemented")
}

func (u *userRepository) UpdateUser(id uuid.UUID, user models.User) (models.User, error) {
	panic("unimplemented")
}

func (u *userRepository) DeleteUser(id uuid.UUID) error {
	panic("unimplemented")
}
