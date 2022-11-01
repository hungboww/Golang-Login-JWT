package model

import "gorm.io/gorm"

type Forum struct {
	gorm.Model
	UserID      int `gorm:"user_id"`
	Description string
	Title       string
	View        int
	Stt         int  `default:"1"`
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
