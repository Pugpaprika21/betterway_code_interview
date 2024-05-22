package repository

import (
	"github.com/Pugpaprika21/question10/dto"
	"gorm.io/gorm"
)

type (
	IReportRepository interface {
		GetAll(filters ...interface{}) []dto.ReportEmployeesResult
	}

	ReportRepository struct {
		query *gorm.DB
	}
)

func NewReportRepository(query *gorm.DB) *ReportRepository {
	return &ReportRepository{
		query: query,
	}
}

func (r *ReportRepository) GetAll(filters ...interface{}) []dto.ReportEmployeesResult {
	var records []dto.ReportEmployeesResult

	query := r.query.Table("employees AS emp").
		Select("emp.id AS emp_id, dep.id AS dep_id, dep.name AS dep_name, emp.name AS emp_name, emp.salary").
		Joins("INNER JOIN departments AS dep ON emp.department_id = dep.id")

	for _, filter := range filters {
		switch v := filter.(type) {
		case dto.ReportEmployeesFilter:
			if v.EmployeeID != "" {
				query = query.Where("emp.id = ?", v.EmployeeID)
			}
			if v.DepartmentID != "" {
				query = query.Where("dep.id = ?", v.DepartmentID)
			}
		}
	}

	query.Order("emp.id DESC").Scan(&records)

	return records
}

/*

SELECT
emp.id AS emp_id,
dep.id AS dep_id,
dep.name AS dep_name,
emp.name AS emp_name,
emp.salary
FROM employees AS emp
INNER JOIN departments AS dep
ON emp.department_id = dep.id
ORDER BY emp.id DESC


*/
