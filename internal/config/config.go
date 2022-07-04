package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Configurations exported
type Configuration struct {
	Debug bool
	Port  int
}

func Config(envPrefix string) Configuration {
	viper.SetEnvPrefix(envPrefix)
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Info().Msg("cannot read config file .env, skipping to env var")
	}
	viper.AutomaticEnv()

	config := Configuration{
		Debug: viper.GetBool("DEBUG"),
		Port:  viper.GetInt("PORT"),
	}

	return config
}
