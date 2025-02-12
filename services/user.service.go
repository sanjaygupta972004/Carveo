package services

import (
	"carveo/config"
	"carveo/models"
	"carveo/notification/email"
	"carveo/repositories"
	"carveo/utils"
	"carveo/utils/auth"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
)

type UserLoginData struct {
	Email            string
	FullName         string
	IsAdmin          bool
	ProfileImage     string
	RefreshToken     string
	AccessToken      string
	UserID           uuid.UUID
	IsMobileVarified bool
	IsEmailVarified  bool
	CreatedAt        string
	UpdatedAt        string
}

type UserService interface {
	RegisterUser(user models.User) (models.User, error)
	VarifyEmail(email string) error
	LoginUser(email, password string) (UserLoginData, error)
	GetUserProfile(id string) (models.User, error)
	UpdateUserProfile(id string, user models.User) (models.User, error)
	DeleteUserProfile(id string) (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) RegisterUser(user models.User) (models.User, error) {
	var reqData models.User
	JwtToken := config.GetConfig().JwtSecret
	baseUrl := os.Getenv("SERVER_BASE_URL")
	if JwtToken == "" && baseUrl == "" {
		return models.User{}, fmt.Errorf("JWT_SECRET environment variable not set")
	}
	// generate email verified token
	emailVerificationToken, err := auth.GenerateEmailVerificationToken(user.Email, JwtToken)
	if err != nil {
		return models.User{}, err
	}

	// hash password
	hashedPassword, err := utils.HashPassword(user.PassWord)
	if err != nil {
		return models.User{}, err

	}
	if hashedPassword != "" {
		reqData = user
		reqData.PassWord = hashedPassword
		reqData.AuthToken = emailVerificationToken
	}
	userData, err := u.userRepository.CreateUser(reqData)
	if err != nil {
		return models.User{}, err
	}

	mailgun := email.NewMailgunService()
	err = mailgun.SendVerificationEmail(userData.Email, emailVerificationToken, userData.FullName, baseUrl)
	if err != nil {
		return models.User{}, err
	}
	return userData, nil
}

func (u *userService) VarifyEmail(email string) error {
	err := u.userRepository.VarifyEmail(email)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) LoginUser(email string, password string) (UserLoginData, error) {
	user, err := u.userRepository.LoginUser(email)
	if err != nil {
		return UserLoginData{}, err
	}
	// generate jwt token and set in cookie
	signedRefreshToken, err := auth.SignJWTForUser(&user)
	if err != nil {
		return UserLoginData{}, err
	}
	signedAccessToken, err := auth.SignJWTAccessTokenForTravelAgency(&user)
	if err != nil {
		return UserLoginData{}, err
	}
	userData := UserLoginData{
		Email:            user.Email,
		FullName:         user.FullName,
		IsAdmin:          user.IsAdmin,
		ProfileImage:     user.ProfileImage,
		RefreshToken:     signedRefreshToken,
		AccessToken:      signedAccessToken,
		UserID:           user.UserID,
		IsMobileVarified: user.IsMobileVerified,
		IsEmailVarified:  user.IsEmailVerified,
		CreatedAt:        user.CreatedAt.String(),
		UpdatedAt:        user.UpdatedAt.String(),
	}
	return userData, nil
}

func (u *userService) GetUserProfile(id string) (models.User, error) {
	uuid, err := utils.IsUUID(id)
	if err != nil {
		return models.User{}, err
	}
	user, err := u.userRepository.GetUserByID(uuid)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userService) UpdateUserProfile(id string, user models.User) (models.User, error) {
	uuid, err := utils.IsUUID(id)
	if err != nil {
		return models.User{}, err
	}
	updatedUser, err := u.userRepository.UpdateUser(uuid, user)
	if err != nil {
		return models.User{}, err
	}
	return updatedUser, nil
}

func (u *userService) DeleteUserProfile(id string) (string, error) {
	uuid, err := utils.IsUUID(id)
	if err != nil {
		return "", err
	}
	res, err := u.userRepository.DeleteUser(uuid)
	if err != nil {
		return "", err
	}
	return res, nil
}
