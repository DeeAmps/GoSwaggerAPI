package controllers

import "net/http"

type LocationController struct{}


func (rc LocationController) GetAllLocations(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (rc LocationController) GetSingleLocation(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {


	}
}
