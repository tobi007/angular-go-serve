package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tobi007/angular-go-serve/bind"
	"github.com/tobi007/angular-go-serve/controllers"
	"github.com/tobi007/angular-go-serve/db"
	"github.com/tobi007/angular-go-serve/middlewares"
	"github.com/tobi007/angular-go-serve/util"

	"math"
	"net/http"
	"os"
	"time"
)

var routerLogger *logrus.Entry
var timeFormat = "02/Jan/2006:15:04:05 -0700"


func NewRouter(bfs *bind.BinaryFileSystem) *gin.Engine {

	routerLogger = util.GetLogger().WithField("ROUTER_INIT", "DB")
	router := gin.Default()
	// Serve the frontend

	router.Use(static.Serve("/", bfs))
	router.Use(CORSMiddleware())
	auth := middlewares.NewAuthMiddleware(db.GetDB())
	authMiddleware, err := auth.Middleware()

	if err != nil {
		routerLogger.Fatal("JWT Error:" + err.Error())
	}

	router.POST("/login", authMiddleware.LoginHandler)

	userController := controllers.NewUserController(db.GetDB())
	router.POST("/signup", userController.Create)

	v1 := router.Group("api/v1")
	v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		articleGroup := v1.Group("article")
		articleGroup.GET("", nil)

	}
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	return router

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Logger is the logrus logger handler
func Logger(log *logrus.Entry) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}


		if len(c.Errors) > 0 {
			log.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			if statusCode > 499 {
				log.Error(msg)
			} else if statusCode > 399 {
				log.Warn(msg)
			} else {
				log.Info(msg)
			}
		}
	}
}
