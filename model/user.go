package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
