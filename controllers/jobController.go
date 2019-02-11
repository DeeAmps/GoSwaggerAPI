package controllers

import (
	"database/sql"
	"net/http"
	"oraclehr.api.com/models"
	"fmt"
	"encoding/json"
	dbDriver "oraclehr.api.com/db"
	"github.com/gorilla/mux"
)

type JobController struct{}


func (rc JobController) GetAllJobs(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		var job models.Job
		jobs := []models.Job{}
		rows := dbDriver.ExecuteRowsQuery(db, query)
		for rows.Next() {
			err := rows.Scan(&job.JobId,
				&job.JobTitle, &job.MinSalary, &job.MaxSalary)
			logFatal(err, db,  rows)
			jobs = append(jobs, job)
		}
		defer rows.Close()
		defer db.Close()
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(jobs)
	}
}

func (rc JobController) GetSingleJob(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		params := mux.Vars(request)
		var job models.Job

		fqry := fmt.Sprintf("%s %s", query, params["id"])
		row := dbDriver.ExecuteRowQuery(db, fqry)
		err := row.Scan(&job.JobId,
			&job.JobTitle, &job.MinSalary, &job.MaxSalary)
		if err == sql.ErrNoRows {
			logFatal(err, db, nil)
			var response = models.Response { false, "No results" }
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(response)
		} else {
			defer db.Close()
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(job)
		}
	}
}
