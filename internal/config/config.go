package config

import (
	"fmt"
	"os"

	"github.com/geniee-ai/geniee-cli/internal/helpers"
	"github.com/spf13/viper"
)

type Config struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

var Cfg Config

func LoadConfig() (*Config, error) {

	homeDir, _ := os.UserHomeDir()
	if !helpers.IsExists(homeDir + "/.geniee/config.json") {
		return nil, fmt.Errorf("could not find config.json")
	}

	viper.SetConfigName("config")        // name of config file (without extension)
	viper.SetConfigType("json")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.geniee") // path to look for the config file in
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	return &Cfg, nil

}
