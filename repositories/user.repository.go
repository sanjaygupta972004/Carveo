package repositories

import (
	"carveo/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	VarifyEmail(email string) error
	RegenerateRefreshAccessToken(accessToken string) (models.User, *gorm.DB, error)
	GenerateResetPasswordToken(email string, resetPasswordToken string) (models.User, error)
	ValidateResetPasswordToken(email, resetPasswordToken string) (string, error)
	UpdatePassword(email string, newPassword string) (string, error)

	LoginUser(email string) (models.User, *gorm.DB, error)
	GetUserByID(id uuid.UUID) (models.User, error)
	UpdateUser(id uuid.UUID, user models.User) (models.User, error)
	DeleteUser(id uuid.UUID) (string, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) UpdatePassword(email string, newPassword string) (string, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", err
	}
	if user.Email == email && user.PassWord != newPassword && newPassword != "" && user.ResetPasswordToken == "" {
		if err := u.db.Model(&user).Where("email = ?", email).Updates(map[string]interface{}{
			"password": newPassword,
		}).Error; err != nil {
			return "", err
		}
	}

	return "Password is successfully updated", nil
}

func (u *userRepository) ValidateResetPasswordToken(email, resetPasswordToken string) (string, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", err
	}
	if user.ResetPasswordToken == resetPasswordToken {
		if err := u.db.Model(&user).Where("email = ?", email).Updates(map[string]interface{}{
			"reset_password_token": "",
		}).Error; err != nil {
			return "", err
		}
	}
	return "Reset password Token is successfully verified", nil
}

func (u *userRepository) GenerateResetPasswordToken(email string, resetPasswordToken string) (models.User, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	if user.Email == email && user.ResetPasswordToken == "" {
		if err := u.db.Model(&user).Where("email = ?", email).Updates(map[string]interface{}{
			"reset_password_token": resetPasswordToken,
		}).Error; err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (u *userRepository) CreateUser(user models.User) (models.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userRepository) VarifyEmail(email string) error {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}
	if user.Email == email {
		if err := u.db.Model(&user).Where("email = ?", email).Updates(map[string]interface{}{
			"is_email_verified": true,
			"auth_token":        "",
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (u *userRepository) LoginUser(email string) (models.User, *gorm.DB, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, u.db, err
	}
	return user, u.db, nil
}

func (u *userRepository) RegenerateRefreshAccessToken(refreshToken string) (models.User, *gorm.DB, error) {
	var user models.User
	err := u.db.Where("access_token = ?", refreshToken).First(&user).Error
	if err != nil {
		return models.User{}, u.db, err
	}
	if user.RefreshToken == refreshToken {
		if err := u.db.Model(&user).Where("refresh_token = ?", refreshToken).Update("refresh_token", "").Error; err != nil {
			return models.User{}, u.db, err
		}
	}
	return user, u.db, nil
}

func (u *userRepository) GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User
	err := u.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userRepository) UpdateUser(id uuid.UUID, user models.User) (models.User, error) {
	var updatedUser models.User
	if err := u.db.Where("user_id = ?", id).First(&updatedUser).Error; err != nil {
		return models.User{}, err
	}

	if err := u.db.Model(&updatedUser).Where("user_id = ?", id).Updates(user).Error; err != nil {
		return models.User{}, err
	}
	return updatedUser, nil
}

func (u *userRepository) DeleteUser(id uuid.UUID) (string, error) {
	var user models.User
	err := u.db.Delete(&user, id).Error

	if err != nil {
		return "Something went wrong while deleting user", err
	}
	return "User deleted successfully", nil
}
