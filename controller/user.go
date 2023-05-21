package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"shop-mall/model"
	"shop-mall/service"
	validatorUtil "shop-mall/validator"
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
		//validator.HandleValidatorError(c, err)
		return
	}

	users, err := service.User().GetUserList(c, model.UserGetListInput{
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

func PassWordLogin(c *gin.Context) {
	in := model.PasswordLoginInput{}

	if err := c.ShouldBind(&in); err != nil {
		validatorUtil.HandleValidatorError(c, err)
		return
	}
	service.User().PasswordLogin(c, in)
}

func Register(c *gin.Context) {
	in := model.CreateUserInput{}

	if err := c.ShouldBind(&in); err != nil {
		validatorUtil.HandleValidatorError(c, err)
		return
	}

	service.User().CreateUser(c, in)
}
