package config

import (
	"os"
)

type Config struct {
	DbName             string
	DbUser             string
	DbPassword         string
	DbHost             string
	DbPort             string
	ServerPort         string
	RecommenderAddress string
}

func (config *Config) ReadConfig() error {
	config.DbName = os.Getenv("POSTGRES_DB")
	config.DbUser = os.Getenv("POSTGRES_USER")
	config.DbPassword = os.Getenv("POSTGRES_PASSWORD")
	config.DbHost = os.Getenv("POSTGRES_HOST")
	config.DbPort = os.Getenv("POSTGRES_PORT")
	config.ServerPort = os.Getenv("SERVER_PORT")
	config.RecommenderAddress = os.Getenv("RECOMMENDER_ADDRESS")

	return nil
}
