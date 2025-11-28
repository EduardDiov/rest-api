package configuration

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	Env         string `yaml:"env" env_required:"true"`
	DatabaseUrl string `yaml:"database_url" env_required:"true"`
	HTTPServer  `yaml:"http_server"`
	Security    `yaml:"security"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type Security struct {
	Key    string `yaml:"key" env-required:"true"`
	Secret string `yaml:"secret" env-required:"true"`
}

func MustLoad() *Configuration {
	configurationPath := "./configuration/configuration.yaml"

	if _, err := os.Stat(configurationPath); os.IsNotExist(err) {
		log.Fatalf("configuration file does not exist: %s", configurationPath)
	}

	var c Configuration

	if err := cleanenv.ReadConfig(configurationPath, &c); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &c
}
