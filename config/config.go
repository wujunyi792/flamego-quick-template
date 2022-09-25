package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wujunyi792/flamego-quick-template/pkg/colorful"
	"github.com/wujunyi792/flamego-quick-template/pkg/fs"
	"gopkg.in/yaml.v3"
	"os"
)

var serveConfig *GlobalConfig

func LoadConfig(configYml string) {
	if !fs.FileExist(configYml) {
		println("cannot find config file")
		os.Exit(1)
	}
	serveConfig = new(GlobalConfig)
	viper.SetConfigFile(configYml)
	err := viper.ReadInConfig()
	if err != nil {
		println("Config Read failed: " + err.Error())
		os.Exit(1)
	}
	err = viper.Unmarshal(serveConfig)
	if err != nil {
		println("Config Unmarshal failed: " + err.Error())
		os.Exit(1)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		println("Config fileHandle changed: ", e.Name)
		_ = viper.ReadInConfig()
		err = viper.Unmarshal(serveConfig)
		if err != nil {
			println("New Config fileHandle Parse Failed: ", e.Name)
			return
		}
	})
	viper.WatchConfig()
}

func GenConfig(configYml string, force bool) error {
	if !fs.FileExist(configYml) || force {
		data, _ := yaml.Marshal(&GlobalConfig{MODE: "debug"})
		err := os.WriteFile(configYml, data, 0644)
		if err != nil {
			return errors.New(colorful.Red("Generate file with error: " + err.Error()))
		}
		println(colorful.Green("config file `config.yaml` generate success in " + configYml))
	} else {
		return errors.New(colorful.Red(configYml + " already exist, use -f to Force coverage"))
	}
	return nil
}

func GetConfig() *GlobalConfig {
	return serveConfig
}
