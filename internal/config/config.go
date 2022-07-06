package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Configurations exported
type Configuration struct {
	Debug    bool
	Port     int
	OAuthApp OAuthApp
	DB       Postgres
}

type OAuthApp struct {
	OAuthRequest  OAuthRequest
	OAuthResponse OAuthResponse
}

type OAuthRequest struct {
	ID     string `json:"client_id"`
	Secret string `json:"client_secret"`
}

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
type Postgres struct {
	Username string
	Password string
	Host     string
	Name     string
}

func Config() Configuration {
	viper.SetEnvPrefix("api")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Info().Msg("cannot read config file .env, skipping to env var")
	}
	viper.AutomaticEnv()

	config := Configuration{
		Debug: viper.GetBool("DEBUG"),
		Port:  viper.GetInt("PORT"),
		OAuthApp: OAuthApp{
			OAuthRequest: OAuthRequest{
				ID:     viper.GetString("OAUTH_APP_ID"),
				Secret: viper.GetString("OAUTH_APP_SECRET"),
			},
		},
		DB: Postgres{
			Username: viper.GetString("USERNAME"),
			Password: viper.GetString("PASSWORD"),
			Host:     viper.GetString("HOST"),
			Name:     viper.GetString("NAME"),
		},
	}

	return config
}
