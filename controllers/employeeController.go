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

type EmployeeController struct{}


func (rc EmployeeController) GetAllEmployees(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		var employee models.Employee
		employees := []models.Employee{}
		rows := dbDriver.ExecuteRowsQuery(db, query)
		for rows.Next() {
			err := rows.Scan(&employee.EmployeeId,
				&employee.FirstName, &employee.LastName, &employee.Email,
				&employee.PhoneNumber, &employee.HireDate, &employee.JobId,
				&employee.Salary, &employee.CommissionPct.Float64,
				&employee.ManagerId.Int64, &employee.DepartmentId.Int64)
			logFatal(err, db,  rows)
			employees = append(employees, employee)
		}
		defer rows.Close()
		defer db.Close()
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(employees)
	}
}

func (rc EmployeeController) GetSingleEmployee(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		params := mux.Vars(request)
		var employee models.Employee
		intParams, _ := strconv.ParseInt(params["id"], 10, 0)
		fqry := fmt.Sprintf("%s %d", query, intParams)
		row := dbDriver.ExecuteRowQuery(db, fqry)
		err := row.Scan(&employee.EmployeeId,
			&employee.FirstName, &employee.LastName, &employee.Email,
			&employee.PhoneNumber, &employee.HireDate, &employee.JobId,
			&employee.Salary, &employee.CommissionPct.Float64,
			&employee.ManagerId.Int64, &employee.DepartmentId.Int64)
		if err == sql.ErrNoRows {
			logFatal(err, db, nil)
			var response = models.Response { false, "No results" }
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(response)
		} else {
			defer db.Close()
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(employee)
		}

	}
}
