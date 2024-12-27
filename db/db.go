package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	conexao := "root:root@tcp(localhost:3306)/go"
	db, err := sql.Open("mysql", conexao)

	if err != nil {
		panic(err)
	}
	return db
}
