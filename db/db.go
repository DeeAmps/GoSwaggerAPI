package db

import (
	"database/sql"
	"fmt"
	_ "gopkg.in/goracle.v2"
	"context"
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
	var ctx = context.Background()
	row := db.QueryRowContext(ctx, query)
	return row
}