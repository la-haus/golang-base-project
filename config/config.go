package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
	"sync"
)

type Config struct {
	ServerPort       string `required:"true" split_words:"true" default:"0.0.0.0"`
	ServerHost       string `required:"true" split_words:"true" default:"3000"`
	DDEnv            string `required:"true" split_words:"true" default:"staging"`
	XApplicationId   string `required:"false" split_words:"true" default:"sample"`
	DDService        string `required:"true" split_words:"true" default:"sample"`
	DDVersion        string `required:"true" split_words:"true" default:"0.0.1"`
	DatadogAgentHost string `required:"true" split_words:"true"`
	DatadogAgentPort string `required:"true" split_words:"true"`
}

var once sync.Once
var c Config

func Environments() Config {
	once.Do(func() {
		if err := envconfig.Process("", &c); err != nil {
			log.Errorf("Error parsing environment vars %#v", err)
		}
	})

	return c
}
