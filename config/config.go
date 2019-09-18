package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tobi007/angular-go-serve/util"
	"log"
	"path/filepath"
)

var config *viper.Viper
var configLogger *logrus.Entry

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	configLogger = util.GetLogger().WithField("CONFIG_INIT", "DB")
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
	}
	configLogger.Info("Config loaded")
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
