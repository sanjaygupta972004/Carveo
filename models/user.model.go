package models

import (
	"carveo/validations"
	"fmt"
	"strings"
	"time"

	"carveo/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID             uuid.UUID      `gorm:"type:uuid;primaryKey;unique; not null; index" json:"userID"`
	FullName           string         `gorm:"not null " json:"fullName"`
	Email              string         `gorm:"unique;not null" json:"email"`
	ProfileImage       string         `gorm:"default:null" json:"profileImage"`
	PassWord           string         `gorm:"not null" json:"password"`
	IsEmailVerified    bool           `gorm:"default:false" json:"isEmailVerified"`
	IsMobileVerified   bool           `gorm:"default:false" json:"isMobileVerified"`
	AuthToken          string         `gorm:"default:null" json:"authToken"`
	ResetPasswordToken string         `gorm:"default:null" json:"resetPasswordToken"`
	RefreshToken       string         `gorm:"default:null" json:"refreshToken"`
	IsAdmin            bool           `gorm:"default:false" json:"isAdmin"`
	MobileNumber       string         `gorm:"default:null" json:"mobileNumber"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	// Relationships
	Car []Car `gorm:"foreignKey:UserID; references:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id := uuid.Must(uuid.NewV4())
	if id != uuid.Nil {
		u.UserID = id
	}
	if err := validations.ValidateUser(validations.User{
		UserID:        u.UserID,
		FullName:      u.FullName,
		Email:         u.Email,
		Password:      u.PassWord,
		ContactNumber: u.MobileNumber,
		PassWord:      u.PassWord,
	}); err != nil {
		return err
	}

	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.PassWord != "" && !strings.HasPrefix(u.PassWord, "$2a$") {
		fmt.Println("Hashing password:", u.PassWord)
		hashedPassword, err := utils.HashPassword(u.PassWord)
		if err != nil {
			return err
		}
		u.PassWord = hashedPassword
	} else {
		fmt.Println("Password already hashed or empty, not hashing again")
	}
	return nil
}

func (User) TableName() string {
	return "Users"
}
