package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	conf "github.com/shumoff/smart-theater-backend/config"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

type application struct {
	config conf.Config
	store  *Store
	server Server
}

func (app *application) start() {
	app.server.Serve()
}

func initApp() (*application, error) {
	var config conf.Config

	if err := config.ReadConfig(); err != nil {
		panicOnErr(err)
	}

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
	server := Server{port: config.Port, store: &store}

	app := application{config: config, server: server, store: &store}

	return &app, nil
}

func main() {
	app, err := initApp()
	panicOnErr(err)

	app.start()
}
