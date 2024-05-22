package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name         string     `gorm:"type:varchar(100);not null"`
	Salary       float64    `gorm:"type:decimal(10,2);not null"`
	DepartmentID int        `gorm:"index"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
}
