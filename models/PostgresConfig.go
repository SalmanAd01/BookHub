package models

import (
	"database/sql"
	"os"
)

// const (
// 	host     = "localhost"
// 	port     = 5433
// 	user     = "postgres"
// 	password = "test123"
// 	dbname   = "Go_Prac1"
// )

func SetupDB() *sql.DB {
	// dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	CheckErr(err)

	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
