package bootstrap

import (
	"fmt"
	"reflect"
	"sail-chat/utils"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func InitTrans() {
	uni := ut.New(zh_Hans_CN.New())
	uni.GetTranslator("zh_Hans_CN")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 获取struct tag 中的 label 为字段名
		v.RegisterTagNameFunc(func(f reflect.StructField) string {
			label := f.Tag.Get("label")
			return label
		})

	}
}

// InitValidation 自定义验证器
func InitValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 手机号验证器
		{
			// 设置自定义验证器配置
			err := v.RegisterValidation("phone", validateMobile)
			if err != nil {
				panic("手机号验证器注册失败" + err.Error())
			}
		}
		{
			if err := v.RegisterValidation("pass", validatePassword); err != nil {
				panic("密码格式错误，应包含至少六位，并包括一个大写字母一个小写字母一个数字 一个特殊字符")
			}
		}
		{
			// 其他验证器
		}
		fmt.Println("初始化验证器成功")
	}
}

// validateMobile 手机号码验证规则
func validateMobile(f validator.FieldLevel) bool {
	mobile := f.Field().String()
	isValid, _ := utils.MatchText(mobile, `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, "phone")
	return isValid
}

// validatePassword 密码验证规则，最少六位，包括一个大写字母一个小写字母一个数字 一个特殊字符
func validatePassword(f validator.FieldLevel) bool {
	pass := f.Field().String()
	isValid, _ := utils.MatchText(pass, `^\S*(?=\S{6,})(?=\S*\d)(?=\S*[A-Z])(?=\S*[a-z])(?=\S*[!@#$%^&*? ])\S*$`, "pass")
	return isValid
}
