package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	GIN_MODE string `mapstructure:"GIN_MODE"`
	PORT     string `mapstructure:"PORT"`

	HOST     string `mapstructure:"HOST"`
	USER     string `mapstructure:"USER"`
	PASSWORD string `mapstructure:"PASSWORD"`
	DB_NAME  string `mapstructure:"DB_NAME"`
	DB_PORT  string `mapstructure:"DB_PORT"`
	SSL_MODE string `mapstructure:"SSL_MODE"`
	TIMEZONE string `mapstructure:"TIMEZONE"`
}

var EnvConfig *Env

// NewEnv loads configuration settings from a file into our ENV structure.
func NewEnv(fileName string) *Env {
	env := Env{}
	viper.SetConfigFile(fileName)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error in reading the .env file %s", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Error in loading the .env file %v", err)
	}

	return &env
}

func InitEnv() {
	EnvConfig = NewEnv(".env")
}
