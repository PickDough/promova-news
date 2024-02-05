package config

import (
	"github.com/spf13/viper"
)

func getViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config.yml")
	v.SetConfigType("yml")
	return v
}

func NewConfig() (*Config, error) {
	v := getViper()
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var config Config
	err = v.Unmarshal(&config)
	return &config, err
}
