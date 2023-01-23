package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Email string `json:"email"`
	UID   string `json:"uid"`
	Token string `json:"token"`
}

var Cfg Config

func init() {
	viper.SetConfigName("config")        // name of config file (without extension)
	viper.SetConfigType("json")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.geniee") // path to look for the config file in
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		fmt.Errorf("error reading config file: %s", err)
	}

}
