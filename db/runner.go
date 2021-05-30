package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	conf "github.com/shumoff/smart-theater-backend/config"
)

func main() {
	config := conf.Config{}

	if err := config.ReadConfig(); err != nil {
		log.Fatal(err)
		return
	}

	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		config.DbUser,
		config.DbPassword,
		config.DbPort,
		config.DbName,
	)

	m, err := migrate.New(
		"file://db/migrations",
		databaseUrl,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
