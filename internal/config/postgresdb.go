package config

import (
	"fmt"
)

type PostgresDBConfig struct {
	DbName   string `mapstructure:"db-name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	LogLevel string `mapstructure:"log-level"`
}

func (cfg *PostgresDBConfig) Validate() error {
	if cfg.DbName == "" {
		return fmt.Errorf("missing postgres db name")
	}

	if cfg.Host == "" {
		return fmt.Errorf("missing postgres host")
	}

	if cfg.Port < 1024 || cfg.Port > 65535 {
		return fmt.Errorf("port number must be between 1024 and 65535 (inclusive)")
	}

	if cfg.User == "" {
		return fmt.Errorf("missing user name")
	}

	if cfg.Password == "" {
		return fmt.Errorf("missing password")
	}

	return nil
}
