package core

import (
	"errors"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

type Config struct {
	Type    string `mapstructure:"type"`
	Name    string `mapstructure:"name"`
	User    string `mapstructure:"user"`
	Pass    string `mapstructure:"pass"`
	Host    string `mapstructure:"host"`
	Port    string `mapstructure:"port"`
	LogMode bool   `mapstructure:"log_mode"`
}

func BuildConfig() (*Config, error) {
	viper.SetEnvPrefix("CORE_DATABASE")
	envFile := os.Getenv("CORE_DATABASE_ENVIRONMENT_FILE")
	if envFile != "" {
		viper.SetConfigFile(envFile)
		err := viper.ReadInConfig()
		if err != nil {
			slog.Error("couldn't read .env file")
			return nil, errors.New("couldn't read .env file")
		}
	} else {
		viper.SetDefault("TYPE", "sqlite")
		viper.SetDefault("NAME", "ordarr.DB")
		viper.SetDefault("USER", "ordarr")
		viper.SetDefault("PASS", "ordarr")
		viper.SetDefault("HOST", "localhost")
		viper.SetDefault("PORT", "5432")
		viper.SetDefault("LOG_MODE", true)
	}
	viper.AutomaticEnv()
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		slog.Error("couldn't unmarshal environment")
		return nil, errors.New("couldn't unmarshal environment")
	}
	return &config, nil
}
