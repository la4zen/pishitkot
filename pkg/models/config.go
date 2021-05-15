package models

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Database database `json:"database,omitempty"`
	Jwt      jwt      `json:"jwt,omitempty"`
}

type jwt struct {
	Key string `json:"key,omitempty"`
}

type database struct {
	Host     string `json:"host,omitempty"`
	Username string `json:"username,omitempty"`
	Sslmode  string `json:"sslmode,omitempty"`
	Dbname   string `json:"dbname,omitempty"`
	Password string
}

func LoadConfig() (*Config, error) {
	godotenv.Load(".env")
	config := &Config{
		Database: database{},
		Jwt:      jwt{},
	}
	v := viper.New()
	v.AddConfigPath("../../config")
	v.SetConfigType("json")
	v.SetConfigName("app")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}
