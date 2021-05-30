package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DbName             string `yaml:"postgres_db"`
	DbUser             string `yaml:"postgres_user"`
	DbPassword         string `yaml:"postgres_password"`
	DbHost             string `yaml:"postgres_host"`
	DbPort             string `yaml:"postgres_port"`
	ServerPort         string `yaml:"server_port"`
	RecommenderAddress string `yaml:"recommender_address"`
}

func (config *Config) ReadConfig() error {
	f, err := os.Open("config.yml")
	if err != nil {
		return fmt.Errorf("could not open config file: %w", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		return fmt.Errorf("could not decode config file: %w", err)
	}

	//TODO SMART-5: start using environmental configuration
	//config.ServerPort = os.Getenv("PORT")
	//config.DbPassword = os.Getenv("POSTGRES_PASSWORD")
	//config.DbUser = os.Getenv("POSTGRES_USER")
	//config.DbName = os.Getenv("POSTGRES_DB")

	return nil
}
