package response

// 统一参数返回

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一服务响应
type Response struct {
	errno  int         // 错误代码 0正确，-1错误，401未登录
	errmsg string      // 错误信息
	data   interface{} // 返回数据
}

const (
	// SuccessCode success re code
	SuccessCode = 0
	// FailCode fail res code
	FailCode = -1
	// FailTokenCode login fail code
	FailTokenCode = 401
)

// ReSuccess 成功返回信息
func ReSuccess(c *gin.Context, msg interface{}, a interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errno":  SuccessCode,
		"errmsg": msg,
		"data":   a,
	})
}

// ReFail 错误返回信息
func ReFail(c *gin.Context, msg interface{}, a interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errno":  FailCode,
		"errmsg": msg,
		"data":   a,
	})
}

// ReFailToken 用户未登录
func ReFailToken(c *gin.Context, msg interface{}, a interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errno":  FailTokenCode,
		"errmsg": msg,
		"data":   a,
	})
}
