package config

type Config struct {
	Database struct {
		Host     string `yaml:"HOST"`
		Port     string `yaml:"PORT"`
		Username string `yaml:"USERNAME"`
		Password string `yaml:"PASSWORD"`
		DBName   string `yaml:"DATABASENAME"`
	} `yaml:"database"`
}
