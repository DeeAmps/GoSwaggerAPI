package models


type Department struct {
	DepartmentId int `json:"DEPARTMENT_ID"`
	DepartmentName string `json:"DEPARTMENT_NAME"`
	ManagerId int `json:"MANAGER_ID"`
	LocationId int `json:"LOCATION_ID"`
}
