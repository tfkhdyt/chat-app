package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"size:10;unique;not null"`
	Email      string `gorm:"size:25;unique;not null"`
	Password   string `gorm:"size:100;not null"`
	PublicKey  string `gorm:"not null" db:"public_key"`
	PrivateKey string `gorm:"not null" db:"private_key"`
}
