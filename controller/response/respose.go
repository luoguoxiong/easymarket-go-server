package response

// 统一参数返回

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// SuccessCode success re code
	SuccessCode = 0
	// FailCode fail res code
	FailCode = -1
	// FailTokenCode login fail code
	FailTokenCode = -2
)

// ReSuccess 成功返回信息
func ReSuccess(c *gin.Context, msg string, a interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": SuccessCode,
		"msg":  msg,
		"data": a,
	})
}

// ReFail 错误返回信息
func ReFail(c *gin.Context, msg string, a interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": FailCode,
		"msg":  msg,
		"data": a,
	})
}

// ReFailToken 用户未登录
func ReFailToken(c *gin.Context, msg string, a interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": FailTokenCode,
		"msg":  msg,
		"data": a,
	})
}
