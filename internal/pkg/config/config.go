package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mrumyantsev/cipher-machines-app/pkg/lib/e"
)

// A Config is the application configuration structure.
type Config struct {
	IsEnableDebugLogs       bool   `envconfig:"ENABLE_DEBUG_LOGS" default:"false"`
	HttpServerListenAddress string `envconfig:"HTTP_SERVER_LISTEN_ADDRESS" default:":8080"`
}

// New creates application configuration.
func New() *Config {
	return new(Config)
}

// Init initializes application configuration.
func (c *Config) Init() error {
	if err := envconfig.Process("", c); err != nil {
		return e.Wrap("could not populate struct with environment variables", err)
	}

	return nil
}
