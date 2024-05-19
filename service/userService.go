package service

import (
	"errors"
	"gorm.io/gorm"
	"se-school-case/initializer"
	"se-school-case/model"
)

// ErrEmailAlreadyExists Custom errors
var ErrEmailAlreadyExists = errors.New("email already exists")

func HandleEmailSubscription(email string) error {
	// Check if email already exists
	exists, err := CheckIfUserExists(email)
	if err != nil {
		return err
	}
	if exists {
		return ErrEmailAlreadyExists
	}

	// Add new email for subscription
	if err := AddUserEmail(email); err != nil {
		return err
	}

	return nil
}

func CheckIfUserExists(email string) (bool, error) {
	var user model.User
	if err := initializer.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func AddUserEmail(email string) error {
	user := model.User{Email: email}
	if err := initializer.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
