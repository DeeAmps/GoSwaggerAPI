package models


type Job struct {
	JobId string `json:"JOB_ID"`
	JobTitle string `json:"JOB_TITLE"`
	MinSalary int64 `json:"MIN_SALARY"`
	MaxSalary int64 `json:"MAX_SALARY"`
}