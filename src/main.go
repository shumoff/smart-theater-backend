package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	conf "github.com/shumoff/smart-theater-backend/config"
	"github.com/shumoff/smart-theater-backend/recommender_grpc"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

type application struct {
	config            conf.Config
	recommenderClient *recommender_grpc.RecommenderClient
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

	recommenderClient := recommender_grpc.RecommenderClient{Address: config.RecommenderAddress}
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
