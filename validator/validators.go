package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"shop-mall/global"
	"strings"
)

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}

	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}

	return rsp
}

func ValidateMobile(field validator.FieldLevel) bool {
	mobile := field.Field().String()

	if ok, _ := regexp.MatchString(`^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`, mobile); !ok {
		return false
	}
	return true
}

func HandleValidatorError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	zap.S().Debug("错误！")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
}
