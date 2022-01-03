package models

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "test123"
	dbname   = "Go_Prac1"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)

	CheckErr(err)

	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
