package initilize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"shop-mall/global"
	"strings"

	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

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
