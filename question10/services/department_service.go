package services

import (
	"net/http"

	"github.com/Pugpaprika21/question10/dto"
	"github.com/Pugpaprika21/question10/repository"
	"github.com/labstack/echo/v4"
)

type departmentService struct {
	repos repository.IDepartmentRepository
}

func NewDepartmentService(repos *repository.DepartmentRepository) *departmentService {
	return &departmentService{
		repos: repos,
	}
}

func (d *departmentService) DepartmentSave(c echo.Context) error {
	var body dto.DepartmentReqBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Respones{
			Message: err.Error(),
		})
	}

	if err := d.repos.Save(body); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Respones{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Message: "save department success..",
	})
}

func (d *departmentService) DepartmentGetAll(c echo.Context) error {
	var data []dto.DepartmentsResp
	departments := d.repos.GetAll()
	for _, department := range departments {
		data = append(data, dto.DepartmentsResp{
			ID:   department.ID,
			Name: department.Name,
		})
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Data: data,
	})
}

func (d *departmentService) DepartmentGetByID(c echo.Context) error {
	departmentID := c.Param("departmentId")
	department := d.repos.GetByID(departmentID)

	if department.ID == 0 {
		return c.JSON(http.StatusOK, dto.Respones{
			Message: "department not found..",
		})
	}

	data := dto.DepartmentsResp{
		ID:   department.ID,
		Name: department.Name,
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Data: data,
	})
}

func (d *departmentService) DepartmentUpdateByID(c echo.Context) error {
	departmentID := c.Param("departmentId")

	var body dto.DepartmentReqBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Respones{
			Message: err.Error(),
		})
	}

	if d.repos.GetByID(departmentID).ID == 0 {
		return c.JSON(http.StatusOK, dto.Respones{
			Message: "department not found..",
		})
	}

	if err := d.repos.UpdateByID(departmentID, body); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Respones{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Message: "update department success..",
	})
}

func (d *departmentService) DepartmentDeleteByID(c echo.Context) error {
	departmentID := c.Param("departmentId")

	if d.repos.GetByID(departmentID).ID == 0 {
		return c.JSON(http.StatusOK, dto.Respones{
			Message: "department not found..",
		})
	}

	if err := d.repos.DeleteByID(departmentID); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Respones{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Respones{
		Message: "delete department success..",
	})
}
