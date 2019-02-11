package models

import (
	"time"
	"database/sql"
)

type Employee struct {
	EmployeeId int `json:"EMPLOYEE_ID"`
	FirstName string `json:"FIRST_NAME"`
	LastName string `json:"LAST_NAME"`
	Email string `json:"EMAIL"`
	PhoneNumber string `json:"PHONE_NUMBER"`
	HireDate time.Time `json:"HIRE_DATE"`
	JobId string `json:"JOB_ID"`
	Salary int64 `json:"SALARY"`
	CommissionPct sql.NullFloat64 `json:"COMMISSION_PCT"`
	ManagerId sql.NullInt64 `json:"MANAGER_ID"`
	DepartmentId sql.NullInt64 `json:"DEPARTMENT_ID"`
}