package configs

import (
	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Address  string `mapstructure:"address"`
	Schema   string `mapstructure:"schema"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Driver   string `mapstructure:"driver"`
}

type Config struct {
	MySql MySQLConfig `mapstructure:"mysql"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("envs")
	vp.SetConfigType("json")
	vp.AddConfigPath("./envs")

	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
