package services

import (
	"net/http"

	"github.com/Pugpaprika21/question10/dto"
	"github.com/Pugpaprika21/question10/repository"
	"github.com/Pugpaprika21/question10/utils"
	"github.com/labstack/echo/v4"
)

type reportServices struct {
	repos repository.IReportRepository
}

func NewReportServices(repos *repository.ReportRepository) *reportServices {
	return &reportServices{
		repos: repos,
	}
}

func (r *reportServices) ReportGetAll(c echo.Context) error {
	var data []dto.ReportEmployeesResp

	filters := dto.ReportEmployeesFilter{
		DepartmentID: c.QueryParam("departmentId"),
		EmployeeID:   c.QueryParam("employeesId"),
	}

	reports := r.repos.GetAll(filters)
	for _, report := range reports {
		data = append(data, dto.ReportEmployeesResp{
			DepartmentName: report.DepartmentName,
			EmployeesName:  report.EmployeesName,
			Salary:         utils.MoneyFormat(report.Salary, false),
		})
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Data: data,
	})
}
