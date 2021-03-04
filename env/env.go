package env

import (
	"github.com/lbryio/lbry.go/v2/extras/errors"

	e "github.com/caarlos0/env"
)

// Config holds the environment configuration used by lighthouse.
type Config struct {
	CoinMineAPIKey string `env:"COINMINE_API_KEY"`
	LbrycrdURL     string `env:"LBRYCRD_CONNECT" envDefault:""`
	SlackHookURL   string `env:"SLACKHOOKURL"`
	SlackChannel   string `env:"SLACKCHANNEL"`
}

// NewWithEnvVars creates an Config from environment variables
func NewWithEnvVars() (*Config, error) {
	cfg := &Config{}
	err := e.Parse(cfg)
	if err != nil {
		return nil, errors.Err(err)
	}

	return cfg, nil
}
