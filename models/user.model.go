package models

import (
	"carveo/utils"
	"carveo/validations"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID             uuid.UUID      `gorm:"type:uuid;primaryKey;unique; not null; index" json:"userID"`
	FullName           string         `gorm:"not null"`
	Email              string         `gorm:"unique;not null"`
	ProfileImage       string         `gorm:"default:null"`
	PassWord           string         `gorm:"not null"`
	IsEmailVerified    bool           `gorm:"default:false"`
	IsMobileVerified   bool           `gorm:"default:false"`
	AuthToken          string         `gorm:"default:null"`
	ResetPasswordToken string         `gorm:"default:null"`
	IsAdmin            bool           `gorm:"default:false"`
	MobileNumber       string         `gorm:"default:null"`
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	// Relationships
	Car []Car `gorm:"foreignKey:UserID; references:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id := uuid.Must(uuid.NewV4())
	if u.UserID != uuid.Nil {
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
