package env

import "github.com/spf13/viper"

func Load() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}