package models

import (
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"`
}
