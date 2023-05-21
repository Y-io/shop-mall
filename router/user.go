package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shop-mall/controller"
)

func InitUserRouter(Router *gin.RouterGroup) {

	zap.S().Debug("初始化用户路由")
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", controller.GetUserList)
		UserRouter.POST("login", controller.PassWordLogin)
		UserRouter.POST("register", controller.Register)
	}
}
