package dto

type EmployeeReqBody struct {
	Name         string  `json:"name"`
	Salary       float64 `json:"salary"`
	DepartmentID int     `json:"departmentId"`
}

type EmployeeResp struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	Salary         float64 `json:"salary"`
	DepartmentID   int     `json:"departmentId"`
	DepartmentName string  `json:"departmentName"`
}
