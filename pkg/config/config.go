package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os/user"
	"path"
)

var cfg *Config
var manager *viper.Viper

type Config struct {
	Server string `json:"server"`
	Token  string `json:"token"`
}

func GetConfig() *Config {
	return cfg
}

func LoadConfig() error {
	manager = viper.New()
	manager.SetConfigName("config")
	info, err := user.Current()
	if err != nil {
		return err
	}
	manager.AddConfigPath(fmt.Sprintf(path.Join(info.HomeDir, ".ko")))
	manager.SetConfigType("yaml")
	if err := manager.ReadInConfig(); err != nil {
		return err
	}
	if err := manager.Unmarshal(&cfg); err != nil {
		return err
	}
	return nil
}

func SetToken(token string) error {
	manager.Set("token", token)
	return manager.WriteConfig()
}
