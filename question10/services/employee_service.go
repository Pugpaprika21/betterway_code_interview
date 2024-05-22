package services

import (
	"net/http"
	"strconv"

	"github.com/Pugpaprika21/question10/dto"
	"github.com/Pugpaprika21/question10/repository"
	"github.com/labstack/echo/v4"
)

type (
	employeeService struct {
		repos repository.IEmployeeRepository
	}
)

func NewEmployeeService(repos *repository.EmployeeRepository) *employeeService {
	return &employeeService{
		repos: repos,
	}
}

func (e *employeeService) EmployeeSave(c echo.Context) error {
	var body dto.EmployeeReqBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Respones{
			Message: err.Error(),
		})
	}

	if err := e.repos.Save(body); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Respones{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Message: "save employee success..",
	})
}

func (e *employeeService) EmployeeGetAll(c echo.Context) error {
	var data []dto.EmployeeResp
	employees := e.repos.GetAll()
	for _, employee := range employees {

		departmentID := strconv.FormatInt(int64(employee.DepartmentID), 10)
		department := e.repos.GetDepartmentByID(departmentID)

		data = append(data, dto.EmployeeResp{
			ID:             employee.ID,
			Name:           employee.Name,
			Salary:         employee.Salary,
			DepartmentID:   employee.DepartmentID,
			DepartmentName: department.Name,
		})
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Data: data,
	})
}

func (e *employeeService) EmployeeGetByID(c echo.Context) error {
	employeeID := c.Param("employeeId")
	employee := e.repos.GetByID(employeeID)

	if employee.ID == 0 {
		return c.JSON(http.StatusOK, dto.Respones{
			Message: "employee not found..",
		})
	}

	departmentID := strconv.FormatInt(int64(employee.DepartmentID), 10)
	department := e.repos.GetDepartmentByID(departmentID)

	data := dto.EmployeeResp{
		ID:             employee.ID,
		Name:           employee.Name,
		Salary:         employee.Salary,
		DepartmentID:   employee.DepartmentID,
		DepartmentName: department.Name,
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Data: data,
	})
}

func (e *employeeService) EmployeeUpdateByID(c echo.Context) error {
	employeeID := c.Param("employeeId")

	var body dto.EmployeeReqBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Respones{
			Message: err.Error(),
		})
	}

	if e.repos.GetByID(employeeID).ID == 0 {
		return c.JSON(http.StatusOK, dto.Respones{
			Message: "employee not found..",
		})
	}

	if err := e.repos.UpdateByID(employeeID, body); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Respones{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Message: "update employee success..",
	})
}

func (e *employeeService) EmployeeDeleteByID(c echo.Context) error {
	employeeID := c.Param("employeeId")

	if e.repos.GetByID(employeeID).ID == 0 {
		return c.JSON(http.StatusOK, dto.Respones{
			Message: "employee not found..",
		})
	}

	if err := e.repos.DeleteByID(employeeID); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Respones{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Message: "delete employee success..",
	})
}
