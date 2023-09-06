package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId  uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(30) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Address string `gorm:"type:varchar(50) not null"`
}
