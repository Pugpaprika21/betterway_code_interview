package dto

type DepartmentReqBody struct {
	Name string `json:"name"`
}

type DepartmentsResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
