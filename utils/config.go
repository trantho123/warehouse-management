package utils

import "github.com/spf13/viper"

type Config struct {
	DBSource            string `mapstructure:"DB_SOURCE"`
	HTTP_SERVER_ADDRESS string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
