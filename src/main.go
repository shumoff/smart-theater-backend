package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	recommender_grpc2 "github.com/shumoff/smart-theater-backend/src/recommender_grpc"

	conf "github.com/shumoff/smart-theater-backend/src/config"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

type application struct {
	config            conf.Config
	recommenderClient *recommender_grpc2.RecommenderClient
	store             *Store
	server            Server
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
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		config.DbName,
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
	)

	db, err := sql.Open("postgres", connString)
	panicOnErr(err)

	err = db.Ping()
	panicOnErr(err)

	recommenderClient := recommender_grpc2.RecommenderClient{Address: config.RecommenderAddress}
	store := Store{db: db}
	server := Server{port: config.ServerPort, store: &store, recommenderClient: &recommenderClient}

	app := application{config: config, server: server, store: &store}

	return &app, nil
}

func main() {
	app, err := initApp()
	panicOnErr(err)

	app.start()
}
