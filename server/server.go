package server

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/sirupsen/logrus"
	"github.com/tobi007/angular-go-serve/bind"
	"github.com/tobi007/angular-go-serve/config"
	"github.com/tobi007/angular-go-serve/util"
)

var serverLogger *logrus.Entry

// Init to setup the router
func Init(bfs *bind.BinaryFileSystem) {
	serverLogger = util.GetLogger().WithField("CONFIG_INIT", "DB")
	config := config.GetConfig()

	r := newRouter()

	// Serve the frontend
	r.Use(static.Serve("/", bfs))

	serverLogger.Info(fmt.Sprintf("Starting Server on %s:%s ", config.GetString("serverAddress"), config.GetString("serverPort")))
	err := r.Run(config.GetString("serverAddress") + ":" + config.GetString("serverPort"))
	if err != nil {
		serverLogger.Info("Error starting server", err)
		panic(err)
	}
}
