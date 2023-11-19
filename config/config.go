package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		OrderConfig   *OrderConfig
		GatewayConfig *GatewayConfig
		UserConfig    *UserConfig
		CartConfig    *CartConfig
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

	Redis struct {
		Url string `env-required:"true" yaml:"redis_url" env:"REDIS_URL"`
	}

	GatewayConfig struct {
		App  `yaml:"app"`
		GRPC `yaml:"grpc"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
	}

	UserConfig struct {
		App  `yaml:"app"`
		GRPC `yaml:"grpc"`
		PG   `yaml:"postgres"`
		Log  `yaml:"logger"`
	}

	OrderConfig struct {
		App  `yaml:"app"`
		GRPC `yaml:"grpc"`
		PG   `yaml:"postgres"`
		Log  `yaml:"logger"`
	}

	CartConfig struct {
		App   `yaml:"app"`
		GRPC  `yaml:"grpc"`
		Redis `yaml:"redis"`
		PG    `yaml:"postgres"`
		Log   `yaml:"logger"`
	}
)

func getServiceFromConfig(path string, service interface{}) error {
	err := cleanenv.ReadConfig(path, service)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(service)
	if err != nil {
		return fmt.Errorf("read env error: %w", err)
	}

	return nil
}

// Creates a new config entity after reading the configuration values
// from the YAML file and environment variables.
func NewConfig() (*Config, error) {
	orderConfig := &OrderConfig{}
	if err := getServiceFromConfig("./config/order.yml", orderConfig); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	gatewayConfig := &GatewayConfig{}
	if err := getServiceFromConfig("./config/gateway.yml", gatewayConfig); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	userConfig := &UserConfig{}
	if err := getServiceFromConfig("./config/user.yml", userConfig); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	cartConfig := &CartConfig{}
	if err := getServiceFromConfig("./config/cart.yml", cartConfig); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	cfg := &Config{
		OrderConfig:   orderConfig,
		GatewayConfig: gatewayConfig,
		UserConfig:    userConfig,
		CartConfig:    cartConfig,
	}

	return cfg, nil
}
