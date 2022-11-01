package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID   uint
	ForumID  uint
	User     User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Forum    []Forum `gorm:"many2many:forum_comments;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Body     string
	Activate bool `default:"True"`
}
