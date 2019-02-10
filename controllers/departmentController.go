package controllers

import (
	"net/http"
)

type DepartmentController struct{}


func (rc DepartmentController) GetAllDepartments(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (rc DepartmentController) GetSingleDepartment(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {


	}
}
