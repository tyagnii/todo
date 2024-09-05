// Provides commandline and environment options parsing
package options

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/caarlos0/env"
)

type Options struct {
	Address    string `env:"ADDRESS" json:"address"`
	DBDSN      string `env:"DATABASE_DSN" json:"database_dsn"`
	ConfigPath string `env:"CONFIG"`
}

func ParseOptions() (Options, error) {
	var cfg Options

	flag.StringVar(&cfg.ConfigPath,
		"c", "",
		"Configuration file path")
	flag.StringVar(&cfg.Address,
		"a", "localhost:8080",
		"Add address and port in format <address>:<port>")
	flag.StringVar(&cfg.DBDSN,
		"d",
		"",
		"Connection string in Postgres format")
	flag.Parse()

	// get env vars
	err := env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}

	if cfg.ConfigPath != "" {
		cfg, err = ReadConfigFile(cfg.ConfigPath)
		if err != nil {
			return cfg, err
		}
	}

	return cfg, nil
}

func ReadConfigFile(path string) (cfg Options, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
