package config

import (
	"log"

	"github.com/caarlos0/env/v7"
)

func NewAgentConfig() AgentConfig {
	var conf AgentConfig
	err := env.Parse(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}

func NewServerConfig() ServerConfig {
	var conf ServerConfig
	err := env.Parse(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
