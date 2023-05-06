package initilize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"shop-mall/global"
	"strings"

	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}

	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}

	return rsp
}

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()

		// 第一个参数是默认语言，后面是可用
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)

		if !ok {
			return fmt.Errorf("uni.GetTranslate(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.Trans)

		case "zh":
			zh_translations.RegisterDefaultTranslations(v, global.Trans)

		default:
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		}

		return
	}
	return
}

func HandleValidatorError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
}
