package base

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	System   SystemConfig   `toml:"system"`
	Database DatabaseConfig `toml:"database"`
	Ldap     LdapConfig     `toml:"ldap"`
}

type SystemConfig struct {
	TokenIssuer      string `toml:"token_issuer" mapstructure:"token_issuer"`
	TokenSecretKey   string `toml:"token_secret_key" mapstructure:"token_secret_key"`
	TokenExpiredTime int    `toml:"token_expired_time" mapstructure:"token_expired_time"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	Port     int    `toml:"port"`
}

type LdapConfig struct {
	Host     string `toml:"host"`
	Password string `toml:"password"`
}

func InitViper(env string) {
	var configFileName = fmt.Sprintf("config.%s", env)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
}

func ReadConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file not found !!!")
		} else {
			panic(err)
		}
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)
	})
	viper.WatchConfig()

	if errUnmarshal := viper.Unmarshal(&Conf); errUnmarshal != nil {
		fmt.Println(errUnmarshal)
	} else {
		fmt.Println(Conf)
	}
}
