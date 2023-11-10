package config

import "os"

type Config struct {
	UserName     string `mapstructure:"UserName"`
	Password     string `mapstructure:"Password"`
	DataBaseName string `mapstructure:"DataBaseName"`
	Host         string `mapstructure:"Host"`
	Port         string `mapstructure:"Port"`
}

var C *Config

func DataBase() {
	C = &Config{
		UserName:     os.Getenv("UserName"),
		Password:     os.Getenv("Password"),
		DataBaseName: os.Getenv("DataBaseName"),
		Host:         os.Getenv("Host"),
		Port:         os.Getenv("PORT"),
	}
}
func Get() *Config {
	return C
}
