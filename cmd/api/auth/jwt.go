package auth

import (
	"simple-douyin/cmd/api/respond"
	"simple-douyin/pkg/config"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/errno"
	"time"

	jwt "github.com/Huang-Yujie/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

var mw *jwt.GinJWTMiddleware

func Init() {
	var err error
	mw, err = jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(config.JWT.Secret),
		Timeout:    config.JWT.Expires,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		TokenLookup:   "query: token, form: token, header: Authorization, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		klog.Fatal(err)
	}
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := GetUserID(c)
		if err != nil {
			respond.Error(c, err)
			c.Abort()
			return
		}
		c.Set(constants.IdentityKey, userID)
		c.Next()
	}
}

func GetUserID(c *gin.Context) (int64, error) {
	claims, err := mw.GetClaimsFromJWT(c)
	if err != nil {
		return 0, errno.UnauthorizedErr.WithMessage(err.Error())
	}
	tempUserID, ok := claims[constants.IdentityKey].(float64)
	userID := int64(tempUserID)
	if !ok || userID <= 0 {
		return 0, errno.UnauthorizedErr.WithMessage("jwt user_id error")
	}
	return userID, nil
}

func GenerateToken(userID int64) (string, error) {
	token, _, err := mw.TokenGenerator(userID)
	return token, err
}
