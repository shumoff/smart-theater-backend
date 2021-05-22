package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/shumoff/smart-theater-backend/utils"
)

type application struct {
	config Config
	store  *Store
	server Server
}

func (app *application) start() {
	app.server.Serve()
}

func initApp() (*application, error) {
	var config Config
	config.ReadConfig()

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
