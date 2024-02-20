package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/rs/zerolog/log"
)

type Config struct {
	HTTPServerConfig
	RiotConfig
	CacheConfig
	GeneralConfig
}

type GeneralConfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
}

type CacheConfig struct {
	Host     string `env:"CACHE_HOST" envDefault:"dragonfly"`
	Port     uint16 `env:"CACHE_PORT" envDefault:"6379"`
	Password string `env:"CACHE_PASSWORD" envDefault:""`
}

type RiotConfig struct {
	ApiKey string `env:"API_KEY" envDefault:"RGAPI-e4a6b1cf-8682-4a3a-bb92-80077bbb5e52"`
}

type HTTPServerConfig struct {
	Port uint16 `env:"HTTP_SERVER_PORT" envDefault:"3003"`
}

func GetServiceConfig() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Err(err).Msg("unable to build config")
	}

	return &cfg
}
