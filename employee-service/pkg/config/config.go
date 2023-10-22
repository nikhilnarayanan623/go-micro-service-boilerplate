package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	EmployeeServiceHost string `mapstructure:"EMPLOYEE_SERVICE_HOST"`
	EmployeeServicePort string `mapstructure:"EMPLOYEE_SERVICE_PORT"`
}

var envs = []string{
	"EMPLOYEE_SERVICE_HOST", "EMPLOYEE_SERVICE_PORT",
}

func LoadConfig() (Config, error) {
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
