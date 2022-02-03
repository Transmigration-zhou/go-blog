package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}
