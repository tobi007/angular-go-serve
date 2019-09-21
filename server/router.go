package server

import (
	"fmt"
	// "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	// "github.com/tobi007/angular-go-serve/bind"
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


func newRouter() *gin.Engine {

	routerLogger = util.GetLogger().WithField("ROUTER_INIT", "DB")
	router := gin.Default()

	router.Use(cORSMiddleware())
	auth := middlewares.NewAuthMiddleware(db.GetDB())
	authMiddleware, err := auth.Middleware()

	if err != nil {
		routerLogger.Fatal("JWT Error:" + err.Error())
	}

	router.POST("/login", authMiddleware.LoginHandler)

	userController := controllers.NewUserController(db.GetDB())
	router.POST("/signup", userController.Create)

	v1 := router.Group("api/v1")
	v1.GET("/health", health)

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

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message":"Heath check Successfully",})
}

func cORSMiddleware() gin.HandlerFunc {
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