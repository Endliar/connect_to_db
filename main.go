package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0611"
	dbname   = "for_golang"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host,
		port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
