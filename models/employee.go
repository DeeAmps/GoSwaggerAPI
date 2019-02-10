package models


type Employee struct {
	EmployeeId int `json:"EMPLOYEE_ID"`
	FirstName string `json:"FIRST_NAME"`
	LastName string `json:"LAST_NAME"`
	Email string `json:"EMAIL"`
	PhoneNumber string `json:"PHONE_NUMBER"`
	HireDate string `json:"HIRE_DATE"`
	JobId string `json:"JOB_ID"`
	Salary int64 `json:"SALARY"`
	CommissionPct int `json:"COMMISSION_PCT"`
	ManagerId int `json:"MANAGER_ID"`
	DepartmentId int `json:"DEPARTMENT_ID"`
}