package initilize

import (
	"github.com/gin-gonic/gin"
	"shop-mall/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	ApiGroup := Router.Group("v1")

	router.InitUserRouter(ApiGroup)

	return Router
}
