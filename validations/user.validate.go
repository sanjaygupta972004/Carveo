package validations

import (
	"errors"
	"regexp"

	"github.com/gofrs/uuid"
)

type User struct {
	UserID        uuid.UUID
	FullName      string
	Email         string
	Password      string
	ContactNumber string
	PassWord      string
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	number := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[!@#~$%^&*()+|_]{1}`)
	uppercase := regexp.MustCompile(`[A-Z]`)

	if !number.MatchString(password) {
		return errors.New("password must contain at least one number")
	}

	if !special.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	if !uppercase.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	return nil
}

func validateE164ContactNumber(number string) error {
	regex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(regex)
	if !re.MatchString(number) {
		return errors.New("enter a valid contact number in E.164 format and it should be contain with + and country code")
	}
	return nil
}

func validateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]{1,64}@[a-zA-Z0-9\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("enter a valid email")
	}
	return nil
}

func ValidateUser(user User) error {
	if user.UserID == uuid.Nil {
		return errors.New("UserID must not be empty")
	}
	if user.FullName == "" {
		return errors.New("full Name must not be empty")
	}
	if user.Email == "" {
		return errors.New("email must not be empty")
	}
	if user.Email != "" {
		if err := validateEmail(user.Email); err != nil {
			return err
		}
	}

	if user.Password == "" {
		return errors.New("password must not be empty")
	}
	if user.Password != "" {
		if err := validatePassword(user.Password); err != nil {
			return err
		}
	}
	if user.ContactNumber == "" {
		return errors.New("contact Number must not be empty")
	}
	if user.ContactNumber != "" {
		if err := validateE164ContactNumber(user.ContactNumber); err != nil {
			return err
		}
	}

	return nil
}
