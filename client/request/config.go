package request

import (
	"net/http"
)

type Config struct {
	Cli *http.Client
}

func (c Config) Ensure() Config {
	if c.Cli == nil {
		c.Cli = &http.Client{}
	}

	return c
}

func (c Config) Verify() {
	if c.Cli == nil {
		panic("Config.Cli must not be empty")
	}
}
