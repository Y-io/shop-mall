package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shop-mall/model"
)

func GetUserList(ctx *gin.Context) {
	zap.S().Debug("获取用户列表页")
}

func PassWordLogin(ctx *gin.Context) {
	in := model.PassWordLoginInput{}

	if err := ctx.ShouldBindJSON(&in); err != nil {
		zap.S().Error()
	}
}
