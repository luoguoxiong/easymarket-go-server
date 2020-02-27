package inits

import (
	_ "easymarket/docs" // swagger
	"easymarket/middleware"
	"easymarket/routers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Logger())
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	APIRouters := Router.Group("")
	routers.DetaultRouter(APIRouters)
	routers.TokenRouter(APIRouters)
	routers.StaticRouter(APIRouters)
	return Router
}
