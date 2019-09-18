package middlewares

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tobi007/angular-go-serve/config"
	"github.com/tobi007/angular-go-serve/models"
	"github.com/tobi007/angular-go-serve/repository"
	"github.com/tobi007/angular-go-serve/repository/dao"
	"github.com/tobi007/angular-go-serve/util"
	"time"
)

func NewAuthMiddleware(db *gorm.DB) *Auth {
	return &Auth{
		dao: dao.NewuserDao(db),
	}
}

type Auth struct {
	dao repository.UserRepo
}

func (a *Auth) Middleware() (*jwt.GinJWTMiddleware, error) {

	config := config.GetConfig()

	identityKey := config.GetString("http.auth.id")

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"email": v.Email,
					"phone": v.Phone,
					"company": v.Company,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Email: claims["email"].(string),
				Phone: claims["phone"].(string),
				Company: claims["company"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userEmail := loginVals.Email
			password := loginVals.Password

			dbUser, err := a.dao.RetrieveByEmail(userEmail)
			if err != nil && dbUser == nil {
				return nil, jwt.ErrFailedAuthentication
			}


			if userEmail == dbUser.Email && util.ComparePassWithHash(dbUser.Password, password) == nil {
				return &models.User{
					Company:  dbUser.Company,
					Email:  dbUser.Email,
					Phone:  dbUser.Phone,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && a.dao.ExistById(v.Email) {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}

