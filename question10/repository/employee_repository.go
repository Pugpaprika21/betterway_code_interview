package repository

import (
	"strings"

	"github.com/Pugpaprika21/question10/dto"
	"github.com/Pugpaprika21/question10/models"
	"gorm.io/gorm"
)

type (
	IEmployeeRepository interface {
		Save(body dto.EmployeeReqBody) error
		GetAll() []models.Employee
		GetByID(id string) models.Employee
		UpdateByID(id string, body dto.EmployeeReqBody) error
		DeleteByID(id string) error
		GetDepartmentByID(id string) models.Department
	}

	EmployeeRepository struct {
		query *gorm.DB
	}
)

func NewEmployeeRepository(query *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		query: query,
	}
}

func (e *EmployeeRepository) Save(body dto.EmployeeReqBody) error {
	data := &models.Employee{
		Name:         strings.TrimSpace(body.Name),
		Salary:       body.Salary,
		DepartmentID: body.DepartmentID,
	}

	tx := e.query.Begin()
	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (e *EmployeeRepository) GetAll() []models.Employee {
	var records []models.Employee
	e.query.Model(&models.Employee{}).Where("deleted_at IS NULL").Order("id DESC").Find(&records)
	return records
}

func (e *EmployeeRepository) GetByID(id string) models.Employee {
	var record models.Employee
	e.query.Model(&models.Employee{}).Where("deleted_at IS NULL AND id = ?", id).First(&record)
	return record
}

func (e *EmployeeRepository) UpdateByID(id string, body dto.EmployeeReqBody) error {
	data := &models.Employee{
		Name:         strings.TrimSpace(body.Name),
		Salary:       body.Salary,
		DepartmentID: body.DepartmentID,
	}

	if err := e.query.Model(&models.Employee{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (e *EmployeeRepository) DeleteByID(id string) error {
	dep := &models.Employee{}
	if err := e.query.Unscoped().Where("id = ?", id).Delete(dep).Error; err != nil {
		return err
	}
	return nil
}

func (e *EmployeeRepository) GetDepartmentByID(id string) models.Department {
	var record models.Department
	e.query.Model(&models.Department{}).Where("deleted_at IS NULL AND id = ?", id).First(&record)
	return record
}
