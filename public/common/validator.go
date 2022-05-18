package common

import (
	"regexp"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ch_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 全局Validate数据校验实列
var Validate *validator.Validate

// 全局翻译器
var Trans ut.Translator

// 初始化Validator数据校验
func InitValidate() {
	chinese := zh.New()
	uni := ut.New(chinese, chinese)
	trans, _ := uni.GetTranslator("zh")
	Trans = trans
	Validate = validator.New()
	_ = ch_translations.RegisterDefaultTranslations(Validate, Trans)
	_ = Validate.RegisterValidation("checkMobile", checkMobile)
	Log.Infof("初始化validator.v10数据校验器完成")
}

func checkMobile(fl validator.FieldLevel) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(fl.Field().String())
}
