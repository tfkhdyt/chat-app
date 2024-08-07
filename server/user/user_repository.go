package user

import (
	"github.com/tfkhdyt/fiber-toolbox/exception"
	"gorm.io/gorm"
)

type UserRepo interface {
	create(user *User) error
	verifyUsernameAvailability(username string) error
	verifyEmailAvailability(email string) error
	FindByEmail(email string) (*User, error)
}

type userRepoPg struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return &userRepoPg{db: db}
}

func (r *userRepoPg) create(user *User) error {
	if err := r.db.Create(user).Error; err != nil {
		return exception.NewBadRequestError("failed to create user", err)
	}

	return nil
}

func (r *userRepoPg) verifyUsernameAvailability(username string) error {
	var user User

	if err := r.db.Select("id").Where("username = ?", username).First(&user).Error; err == nil {
		return exception.NewBadRequestError("username is already taken", nil)

	}

	return nil
}

func (r *userRepoPg) verifyEmailAvailability(email string) error {
	var user User

	if err := r.db.Select("id").Where("email = ?", email).First(&user).Error; err == nil {
		return exception.NewBadRequestError("email is already taken", nil)
	}

	return nil
}

func (r *userRepoPg) FindByEmail(email string) (*User, error) {
	var user User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, exception.NewNotFoundError("user not found", err)
	}

	return &user, nil
}
