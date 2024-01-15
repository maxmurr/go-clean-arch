package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
)

func GetConfig() Config {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v \n", err))
	}

	return Config{
		App: App{
			Port: viper.GetInt("APP_PORT"),
		},
		Db: Db{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
		},
	}
}
