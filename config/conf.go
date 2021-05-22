package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/shumoff/smart-theater-backend/utils"
)

type Config struct {
	DbName     string `yaml:"postgres_db"`
	DbUser     string `yaml:"postgres_user"`
	DbPassword string `yaml:"postgres_password"`
	Port       string `yaml:"port"`
}

func (config *Config) ReadConfig() {
	f, err := os.Open("config.yml")
	utils.PanicOnErr(err)
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	utils.PanicOnErr(err)
	//config.DbName = os.Getenv("POSTGRES_DB")
	//config.DbUser = os.Getenv("POSTGRES_USER")
	//config.DbPassword = os.Getenv("POSTGRES_PASSWORD")
	//config.Port = os.Getenv("PORT")
}
