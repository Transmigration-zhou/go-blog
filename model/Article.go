package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category    Category `gorm:"foreignKey:Cid"`
	Title       string   `gorm:"type:varchar(20);not null" json:"title"`
	Cid         uint     `gorm:"type:uint;not null;" json:"cid"`
	Description string   `gorm:"type:varchar(100);not null" json:"description"`
	Content     string   `gorm:"type:longtext" json:"content"`
	Image       string   `gorm:"type:varchar(100)" json:"image"`
}
