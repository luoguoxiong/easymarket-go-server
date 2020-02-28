package inits

import (
	"easymarket/middleware"
	"easymarket/routers"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Logger())
	APIRouters := Router.Group("")
	routers.DetaultRouter(APIRouters)
	routers.TokenRouter(APIRouters)
	routers.StaticRouter(APIRouters)
	return Router
}
