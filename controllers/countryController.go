package controllers


import (
	"net/http"
)

type CountryController struct{}


func (rc CountryController) GetAllCountries(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (rc CountryController) GetSingleCountry(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {


	}
}
