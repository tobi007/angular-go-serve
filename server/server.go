package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tobi007/angular-go-serve/bind"
	"github.com/tobi007/angular-go-serve/config"
	"github.com/tobi007/angular-go-serve/util"
)

var serverLogger *logrus.Entry


func Init(bfs *bind.BinaryFileSystem) {
	serverLogger = util.GetLogger().WithField("CONFIG_INIT", "DB")
	config := config.GetConfig()
	r := NewRouter(bfs)
	serverLogger.Info(fmt.Sprintf("Starting Server on %s:%s ", config.GetString("serverAddress"), config.GetString("serverPort")))
	err := r.Run(config.GetString("serverAddress") +":" + config.GetString("serverPort"))
	if err != nil {
		serverLogger.Info("Error starting server", err)
		panic(err)
	}
}