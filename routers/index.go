package routers

import (
	"easymarket/controller/api"
	"easymarket/middleware"

	"github.com/gin-gonic/gin"
)

// DetaultRouter 不需要验证登录
func DetaultRouter(Router *gin.RouterGroup) {
	v1 := Router.Group("/api")
	{
		v1.GET("/", api.GetHome)
		v1.GET("/topic/list", api.GetTopicList)
		v1.GET("/topic/detail", api.GetTopicDetail)
	}
}

// TokenRouter 需要验证登录
func TokenRouter(Router *gin.RouterGroup) {
	v2 := Router.Group("/api", middleware.Token())
	{
		v2.GET("/login", api.GetHome)
	}
}

// StaticRouter 静态文件路由
func StaticRouter(Router *gin.RouterGroup) {
	v3 := Router.Group("/")
	v3.Static("/static", "./static")
}
