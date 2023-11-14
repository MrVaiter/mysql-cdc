package cdc

import (
	"github.com/rs/zerolog/log"

	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Flavor   string        `envconfig:"MYSQL_FLAVOR" default:"mysql"`
	Host     string        `envconfig:"MYSQL_HOST" default:"127.0.0.1"`
	Port     int           `envconfig:"MYSQL_PORT" default:"3306"`
	User     string        `envconfig:"MYSQL_USER" default:"root"`
	Password string        `envconfig:"MYSQL_PASSWORD" default:"d7Vn4xYggr"`
	Timeout  time.Duration `envconfig:"TIMEOUT" default:"5s"`
	Sleep    time.Duration `envconfig:"TIMEOUT" default:"200ms"`
	DB       string        `envconfig:"MYSQL_DB" default:"mydata"`
}

func ReadConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	return &config, nil
}
