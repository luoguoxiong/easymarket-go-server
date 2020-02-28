package middleware

import (
	"easymarket/controller/response"

	"github.com/gin-gonic/gin"
)

// Token 验证登录中间件
func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			response.ReFailToken(c, nil, "请选登录！")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
