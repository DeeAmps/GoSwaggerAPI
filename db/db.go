package db

import (
	"database/sql"
	"fmt"
	_ "gopkg.in/goracle.v2"
)

func DBConnection(connectionString string) *sql.DB {
	db, err := sql.Open("goracle", connectionString)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return db
}


func ExecuteRowsQuery(db *sql.DB, query string) *sql.Rows {
	rows , err := db.Query(query)
	if err != nil {
		panic(err)
	}
	return rows
}

func ExecuteRowQuery(db *sql.DB, query string) *sql.Row {
	row := db.QueryRow(query)
	return row
}


func ExecuteDataManipulationQuery(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	return err
}