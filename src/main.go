package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

type App struct {
	config Config
	store  Store
	server Server
}

var app App

func initApp(app *App) error {
	var config Config
	config.readConfig()

	connString := fmt.Sprintf(
		"dbname=%s user=%s password=%s sslmode=disable",
		config.DbName,
		config.DbUser,
		config.DbPassword,
	)
	db, err := sql.Open("postgres", connString)
	panicOnErr(err)

	err = db.Ping()
	panicOnErr(err)

	store := Store{db: db}
	server := Server{port: config.Port}

	app.config = config
	app.store = store
	app.server = server

	return nil
}

func main() {
	err := initApp(&app)
	panicOnErr(err)
	app.server.serve()
}
