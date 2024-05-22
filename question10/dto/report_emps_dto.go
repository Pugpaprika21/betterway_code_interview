package dto

type ReportEmployeesResult struct {
	EmployeeID     uint    `gorm:"column:emp_id"`
	DepartmentID   uint    `gorm:"column:dep_id"`
	DepartmentName string  `gorm:"column:dep_name"`
	EmployeesName  string  `gorm:"column:emp_name"`
	Salary         float64 `gorm:"column:salary"`
}

type ReportEmployeesFilter struct {
	EmployeeID     string
	DepartmentID   string
	DepartmentName string
	EmployeesName  string
	Salary         float64
}

type ReportEmployeesResp struct {
	EmployeeID     uint    `json:"employeeId,omitempty"`
	DepartmentID   uint    `json:"departmentId,omitempty"`
	DepartmentName string  `json:"departmentName"`
	EmployeesName  string  `json:"employeesName"`
	SalaryBase     float64 `json:"salaryBase,omitempty"`
	Salary         string  `json:"Salary"`
}
