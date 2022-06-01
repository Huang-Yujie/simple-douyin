package respond

import (
	"net/http"
	"simple-douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, rawErr error) {
	err := errno.ConvertErr(rawErr)
	c.JSON(err.StatusCode(), err)
}

func Send(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, Success)
}
