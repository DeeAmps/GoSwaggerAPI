package controllers


import (
	"net/http"
)

type EmployeeController struct{}


func (rc EmployeeController) GetAllEmployees(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (rc EmployeeController) GetSingleEmployee(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {


	}
}
