package repository

import (
	"strings"

	"github.com/Pugpaprika21/question10/dto"
	"github.com/Pugpaprika21/question10/models"
	"gorm.io/gorm"
)

type (
	IDepartmentRepository interface {
		Save(body dto.DepartmentReqBody) error
		GetAll() []models.Department
		GetByID(id string) models.Department
		UpdateByID(id string, body dto.DepartmentReqBody) error
		DeleteByID(id string) error
	}

	DepartmentRepository struct {
		query *gorm.DB
	}
)

func NewDepartmentRepository(query *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{
		query: query,
	}
}

func (d *DepartmentRepository) Save(body dto.DepartmentReqBody) error {
	data := &models.Department{
		Name: strings.TrimSpace(body.Name),
	}

	tx := d.query.Begin()
	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (d *DepartmentRepository) GetAll() []models.Department {
	var records []models.Department
	d.query.Model(&models.Department{}).Where("deleted_at IS NULL").Order("id DESC").Find(&records)
	return records
}

func (d *DepartmentRepository) GetByID(id string) models.Department {
	var record models.Department
	d.query.Model(&models.Department{}).Where("deleted_at IS NULL AND id = ?", id).First(&record)
	return record
}

func (d *DepartmentRepository) UpdateByID(id string, body dto.DepartmentReqBody) error {
	data := &models.Department{
		Name: strings.TrimSpace(body.Name),
	}

	if err := d.query.Model(&models.Department{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (d *DepartmentRepository) DeleteByID(id string) error {
	dep := &models.Department{}
	if err := d.query.Unscoped().Where("id = ?", id).Delete(dep).Error; err != nil {
		return err
	}
	return nil
}
