package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:10;unique;not null"`
	Email    string `gorm:"size:25;unique;not null"`
	Password string `gorm:"size:100;not null"`
}
