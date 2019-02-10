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
	fmt.Println("DB connection Successful")
	return db
}


func ExecuteQuery(db *sql.DB, query string) *sql.Rows{
	rows , err := db.Query(query)
	if err != nil {
		panic(err)
	}
	return rows
}