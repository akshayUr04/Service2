package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	M2PORT         string `mapstructure:"M2PORT"`
	UserServiceUrl string `mapstructure:"USER_SERVICE_URL"`
}

var envs = []string{"M2PORT", "USER_SERVICE_URL"}

func LoadConfig() (cfg Config, err error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
