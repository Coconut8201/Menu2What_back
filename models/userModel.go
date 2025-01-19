package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"type:varchar(36);uniqueIndex;not null"`
	Username string `json:"username" gorm:"type:varchar(50);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}
