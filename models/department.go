package models

import "database/sql"

type Department struct {
	DepartmentId int `json:"DEPARTMENT_ID"`
	DepartmentName string `json:"DEPARTMENT_NAME"`
	ManagerId sql.NullInt64 `json:"MANAGER_ID"`
	LocationId int `json:"LOCATION_ID"`
}
