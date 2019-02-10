package controllers

import "net/http"

type JobController struct{}


func (rc JobController) GetAllJobs(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (rc JobController) GetSingleJob(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {


	}
}
