package config

import (
	"github.com/spf13/viper"
)

var requiredProps = [2]string{"client_id", "client_secret"}

func loadConfig() map[string]string {
	config := make(map[string]string, 2)

	viper.SetConfigType("json")
	viper.SetConfigFile("dev.json")
	viper.ReadInConfig()

	for _, prop := range requiredProps {
		if viper.IsSet(prop) {
			config[prop] = viper.GetString(prop)
		} else {
			panic("wrong config")
		}
	}

	return config
}
