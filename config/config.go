package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		OrderConfig   *OrderConfig
		GatewayConfig *GatewayConfig
	}

	OrderConfig struct {
		App  `yaml:"app"`
		GRPC `yaml:"grpc"`
		PG   `yaml:"postgres"`
		Log  `yaml:"logger"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	GRPC struct {
		Port uint32 `env-required:"true" yaml:"port" env:"GRPC_PORT"`
		Host string `env-required:"true" yaml:"host" env:"HOST"`
	}

	PG struct {
		Dialect string `env-requered:"true" yaml:"dialect" env:"DIALECT"`
		URL     string `env-required:"true" yaml:"pg_url" env:"PG_URL"`
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	HTTP struct {
		Port uint32 `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		Host string `env-required:"true" yaml:"host" env:"HOST"`
	}

	GatewayConfig struct {
		App  `yaml:"app"`
		GRPC `yaml:"grpc"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
	}
)

// Creates a new config entity after reading the configuration values
// from the YAML file and environment variables.
func NewConfig() (*Config, error) {
	orderConfig := &OrderConfig{}

	err := cleanenv.ReadConfig("./config/order.yml", orderConfig)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(orderConfig)
	if err != nil {
		return nil, err
	}

	gatewayConfig := &GatewayConfig{}

	err = cleanenv.ReadConfig("./config/gateway.yml", gatewayConfig)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(gatewayConfig)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		OrderConfig:   orderConfig,
		GatewayConfig: gatewayConfig,
	}

	return cfg, nil
}
