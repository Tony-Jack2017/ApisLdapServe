package base

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper(env string) {
	var configFileName = fmt.Sprintf("config.%s", env)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("TOML")
	viper.AddConfigPath("./config")
}

func ReadConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			panic("")
		}
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)
	})
	viper.WatchConfig()
}
