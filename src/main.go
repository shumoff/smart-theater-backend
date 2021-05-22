package main

import (
	//...
	// The libn/pq driver is used for postgres
	"database/sql"
	_ "github.com/lib/pq"
	//...
)

func main() {
	// ...
	connString := "dbname=test_db user=dostavisor_user password=password sslmode=disable"

	db, err := sql.Open("postgres", connString)
	panicOnErr(err)

	err = db.Ping()
	panicOnErr(err)

	initStore(&dbStore{db: db})
}

// panicOnErr panics err is not nil.
func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
