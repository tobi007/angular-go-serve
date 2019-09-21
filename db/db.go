package db

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/sirupsen/logrus"
	"github.com/tobi007/angular-go-serve/config"
	// "github.com/tobi007/angular-go-serve/models"
	"github.com/tobi007/angular-go-serve/util"
)

var db *gorm.DB
var err error
var dgLogger *logrus.Entry

// ConnectSQL ...
func Init() {
	dgLogger = util.GetLogger().WithField("DB_INIT", "DB")
	c := config.GetConfig()
	dbSource := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%d?database=%s",
		c.Get("DB_USERNAME"),
		c.Get("DB_PASSWORD"),
		c.Get("DB_IP"),
		c.Get("DB_PORT"),
		c.Get("DB_NAME"),
	)

	db, err = gorm.Open("mssql", dbSource)
	if err != nil {
		dgLogger.Info("Failed to connect to database: ", err)
		//dgLogger.Fatal(err)
	}

	gorm.AddNamingStrategy(&gorm.NamingStrategy{
		Column: func(name string) string {
			return strcase.ToLowerCamel(name)
		},
	})

	//db.AutoMigrate(&models.User{}, models.User{})

}

func GetDB() *gorm.DB {
	return db
}