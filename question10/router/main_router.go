package router

import (
	"github.com/Pugpaprika21/question10/repository"
	"github.com/Pugpaprika21/question10/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EchoSetupRouter(e *echo.Echo, dbc *gorm.DB) {
	g := e.Group("/api")

	employeeRepos := repository.NewEmployeeRepository(dbc)
	employee := services.NewEmployeeService(employeeRepos)
	{
		g.POST("/employee", employee.EmployeeSave)
		g.GET("/employees", employee.EmployeeGetAll)
		g.GET("/employee/:employeeId", employee.EmployeeGetByID)
		g.PUT("/employee/:employeeId", employee.EmployeeUpdateByID)
		g.DELETE("/employee/:employeeId", employee.EmployeeDeleteByID)
	}

	departmentRepos := repository.NewDepartmentRepository(dbc)
	department := services.NewDepartmentService(departmentRepos)
	{
		g.POST("/department", department.DepartmentSave)
		g.GET("/departments", department.DepartmentGetAll)
		g.GET("/department/:departmentId", department.DepartmentGetByID)
		g.PUT("/department/:departmentId", department.DepartmentUpdateByID)
		g.DELETE("/department/:departmentId", department.DepartmentDeleteByID)
	}

	repostRepos := repository.NewReportRepository(dbc)
	repost := services.NewReportServices(repostRepos)
	{
		g.GET("/reports", repost.ReportGetAll)
	}
}

// /reports/employees/:employeeId // params http://localhost/api/reports/employees/{empId}/departments/{depId}
