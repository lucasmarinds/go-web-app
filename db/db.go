package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaComBanco() *sql.DB {
	connStr := "user=postgres dbname=marin_loja password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	return db
}
