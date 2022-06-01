package auth

import (
	"simple-douyin/cmd/api/respond"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/errno"
	"time"

	jwt "github.com/Huang-Yujie/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

var config = &jwt.GinJWTMiddleware{
	Key:        []byte(constants.SecretKey),
	Timeout:    time.Hour * 24,
	MaxRefresh: time.Hour,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(int64); ok {
			return jwt.MapClaims{
				constants.IdentityKey: v,
			}
		}
		return jwt.MapClaims{}
	},
	// Authenticator: func(c *gin.Context) (interface{}, error) {
	// 	var loginVar handlers.UserParam
	// 	if err := c.ShouldBind(&loginVar); err != nil {
	// 		return "", jwt.ErrMissingLoginValues
	// 	}

	// 	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
	// 		return "", jwt.ErrMissingLoginValues
	// 	}

	// 	return rpc.CheckUser(context.Background(), &userproto.CheckUserReq{UserName: loginVar.UserName, Password: loginVar.PassWord})
	// },
	TokenLookup:   "query: token, form: token, header: Authorization, cookie: jwt",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
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
	mw, err := jwt.New(config)
	if err != nil {
		klog.Fatal(err)
	}
	claims, err := mw.GetClaimsFromJWT(c)
	if err != nil {
		return 0, errno.UnauthorizedErr.WithMessage(err.Error())
	}
	userID, ok := claims[constants.IdentityKey].(int64)
	if !ok || userID <= 0 {
		return 0, errno.UnauthorizedErr.WithMessage("jwt user_id error")
	}
	return userID, nil
}

func GenerateToken(userID int64) (string, error) {
	mw, err := jwt.New(config)
	if err != nil {
		klog.Fatal(err)
	}
	token, _, err := mw.TokenGenerator(userID)
	return token, err
}
