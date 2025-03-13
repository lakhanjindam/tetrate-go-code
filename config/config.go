package config

import (
	"fmt"

	goenv "github.com/caitlinelfring/go-env-default"
	"github.com/spf13/viper"
)

type APIConfig struct {
	Port         string   `yaml:"port"`
	Env          string   `yaml:"env"`
	Region       string   `yaml:"region"`
	AllowOrigins []string `yaml:"allowOrigins"`
}

var Config = &APIConfig{}

// LoadConfig loads the configuration from the config file and environment variables
func LoadConfig(filePath string) error {

	// Load environment variables into viper
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(filePath)

	// environment variables
	env := goenv.GetDefault("ENV", "local")

	// Load the configuration file
	viper.SetConfigName(env)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(Config); err != nil {
		return fmt.Errorf("unable to decode viper config into struct: %w", err)
	}

	return nil

}
