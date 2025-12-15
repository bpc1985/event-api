package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null" binding:"required"`
	Password string `gorm:"not null" binding:"required" json:"-"`
}

func (u *User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	result := db.DB.Create(u)
	return result.Error
}

func (u *User) ValidateCredentials() error {
	var retrievedUser User
	result := db.DB.Where("email = ?", u.Email).First(&retrievedUser)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("credentials are invalid")
		}
		return result.Error
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedUser.Password)
	if !passwordIsValid {
		return errors.New("credentials are invalid")
	}

	u.ID = retrievedUser.ID
	return nil
}
