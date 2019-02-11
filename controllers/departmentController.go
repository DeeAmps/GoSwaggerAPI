package controllers

import (
	"database/sql"
	"net/http"
	"oraclehr.api.com/models"
	"fmt"
	"encoding/json"
	dbDriver "oraclehr.api.com/db"
	"github.com/gorilla/mux"
	"strconv"
)

type DepartmentController struct{}


func (rc DepartmentController) GetAllDepartments(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		var department models.Department
		departments := []models.Department{}
		rows := dbDriver.ExecuteRowsQuery(db, query)
		for rows.Next() {
			err := rows.Scan(&department.DepartmentId,
				&department.DepartmentName, &department.ManagerId.Int64, &department.LocationId)
			logFatal(err, db,  rows)
			departments = append(departments, department)
		}
		defer rows.Close()
		defer db.Close()
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(departments)
	}
}

func (rc DepartmentController) GetSingleDepartment(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		params := mux.Vars(request)
		var department models.Department
		intParams, _ := strconv.ParseInt(params["id"], 10, 0)

		fqry := fmt.Sprintf("%s %d", query, intParams)
		row := dbDriver.ExecuteRowQuery(db, fqry)
		err := row.Scan(&department.DepartmentId,
			&department.DepartmentName, &department.ManagerId.Int64, &department.LocationId)
		if err == sql.ErrNoRows {
			logFatal(err, db, nil)
			var response = models.Response { false, "No results" }
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(response)
		} else {
			defer db.Close()
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(department)
		}
	}
}
