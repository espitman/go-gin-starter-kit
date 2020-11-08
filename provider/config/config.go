package config

import "github.com/spf13/viper"

func init() {
	readConfig()
}

func readConfig() {
	viper.SetConfigName("default")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	_ = viper.ReadInConfig()
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
