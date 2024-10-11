package utils

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhT "github.com/go-playground/validator/v10/translations/zh"
)

// RemoveTopStruct 移除打印的错误信息中的结构体包前缀
func RemoveTopStruct(fields map[string]string) string {
	var prefix strings.Builder
	for _, v := range fields {
		prefix.WriteString(v)
		prefix.WriteString(";")
	}
	return prefix.String()
}

// HandleValidatorPhone 校验手机号
func HandleValidatorPhone(data interface{}) (msg string) {
	msg = ValidatorError(data, func(v *validator.Validate, trans ut.Translator) {
		// 设置自定义翻译器配置
		_ = v.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
			return ut.Add("phone", "{0} 非法的手机号", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("phone", fe.Field())
			return t
		})
	})
	return msg
}

// ValidatorError 返回校验错误消息
func ValidatorError(data interface{}, fn func(v *validator.Validate, s ut.Translator)) (msg string) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		uni := ut.New(zh_Hans_CN.New())
		trans, _ := uni.GetTranslator("zh_Hans_CN")
		// 注册默认翻译
		_ = zhT.RegisterDefaultTranslations(v, trans)
		// 设置自定义翻译器配置
		fn(v, trans)
		if err := v.Struct(data); err != nil {
			errs, _ := err.(validator.ValidationErrors)
			msg := RemoveTopStruct(errs.Translate(trans))
			return msg
		}
	}
	return ""
}

// ValidatorDefault 默认校验
func ValidatorDefault(data interface{}) (msg string) {
	return ValidatorError(data, func(v *validator.Validate, trans ut.Translator) {})
}
