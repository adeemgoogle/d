package config

import "github.com/spf13/viper"

type Config struct {
	DB string `mapstructure:"DB"`
}

func LoadConfig() (*Config, error) {
	config := new(Config)

	v := viper.New()
	v.AutomaticEnv()

	err := v.BindEnv("DB")
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
