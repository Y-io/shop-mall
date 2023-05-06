package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"shop-mall/model"
	"shop-mall/service"
	"shop-mall/validator"
	"strconv"
)

func GetUserList(c *gin.Context) {
	zap.S().Debug("获取用户列表页")
	page := c.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	pageSize := c.DefaultQuery("pageSize", "10")
	pageSizeInt, _ := strconv.Atoi(pageSize)
	in := model.UserGetListInput{}

	if err := c.ShouldBind(&in); err != nil {
		validator.HandleValidatorError(c, err)
		return
	}

	users, err := service.User().GetUserList(model.UserGetListInput{
		Page:     int32(pageInt),
		PageSize: int32(pageSizeInt),
	})

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func PassWordLogin(ctx *gin.Context) {
	in := model.PassWordLoginInput{}

	if err := ctx.ShouldBindJSON(&in); err != nil {
		validator.HandleValidatorError(ctx, err)
	}

}
