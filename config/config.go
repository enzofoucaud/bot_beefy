package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Discord *Discord
	}

	Discord struct {
		Token string
	}
)

var (
	once        sync.Once
	cfgInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		cfgInstance = &Config{
			Discord: &Discord{
				Token: viper.GetString("discord.token"),
			},
		}
	})

	return cfgInstance
}
