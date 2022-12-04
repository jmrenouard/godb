package main

import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ExecTestDB(dbname string) {
	var db *sqlx.DB
	var err error
	// exactly the same as the built-in
	db, err = sqlx.Open("sqlite3", dbname)
	if err != nil {
		log.Fatalln(err)
	}

	// force a connection and test that it worked
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	schema := `CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer);`

	// execute a query on the server
	result, err := db.Exec(schema)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("result: %v\n", result)

	// or, you can use MustExec, which panics on error
	cityState := `INSERT INTO place (country, telcode) VALUES (?, ?)`
	countryCity := `INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)`
	db.MustExec(cityState, "Hong Kong", 852)
	db.MustExec(cityState, "Singapore", 65)
	db.MustExec(countryCity, "South Africa", "Johannesburg", 27)
}
