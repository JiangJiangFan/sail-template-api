package res

import (
	"net/http"
	"sail-chat/types"

	"github.com/gin-gonic/gin"
)

type Wrapper struct {
	ctx *gin.Context
}

type API struct {
	Code    int    			`json:"code"`
	Msg 		string 			`json:"msg"`
	Data    interface{} `json:"data"`
	Meta    *types.Meta  `json:"meta"`
}

func New(code int, msg string, data interface{}, meta *types.Meta) API {
	// 判断meta是否为空
	return API{
		Code:    code,
		Msg:     msg,
		Data:    data,
		Meta:    meta,
	}
}

// Http 请求
func Http(ctx *gin.Context) *Wrapper {
	return &Wrapper{ctx: ctx}
}

func (wrapper *Wrapper) Success(data interface{}, meta *types.Meta) {
	wrapper.ctx.JSON(http.StatusOK, New(200,"success",data, meta))
}

func (wrapper *Wrapper) SuccessOnly(data interface{}) {
	wrapper.ctx.JSON(http.StatusOK, New(200,"success",data, nil))
}

func (wrapper *Wrapper) Error(code int, msg string) {
	wrapper.ctx.Set("error", code)
	wrapper.ctx.AbortWithStatusJSON(http.StatusOK, New(code, msg, nil, nil))
}

// ErrorParam 服务级错误码
func (wrapper *Wrapper) ErrorParam() {
	wrapper.Error(10010, "参数解析错误")
}

func (wrapper *Wrapper) ErrorSignParam() {
	wrapper.Error(10011, "签名参数有误")
}

func (wrapper *Wrapper) ErrorQuery() {
	wrapper.Error(10012, "SQL 执行错误")
}

func (wrapper *Wrapper) ErrorContentType() {
	wrapper.Error(10013, "contentType 类型错误")
}

// ErrorLoginPass 模块级错误码
func (wrapper *Wrapper) ErrorLoginPass() {
	wrapper.Error(10001, "密码错误")
}

func (wrapper *Wrapper) ErrorLoginNot() {
	wrapper.Error(10002, "用户不存在")
}
func (wrapper *Wrapper) ErrorLoginExist() {
	wrapper.Error(10003, "用户已存在")
}
func (wrapper *Wrapper) ErrorLoginToken() {
	wrapper.Error(10004, "token 无效")
}
func (wrapper *Wrapper) ErrorValidator(msg string) {
	wrapper.Error(10005, msg)
}
func (wrapper *Wrapper) ErrorParse() {
	wrapper.Error(10004, "解析结果错误")
}

func (wrapper *Wrapper) ErrorTrans(msg string) {
	wrapper.Error(http.StatusCreated, msg)
}

func (wrapper *Wrapper) ErrorTransFail(msg string) {
	wrapper.Error(10004, msg)
}

func (wrapper *Wrapper) ErrorSQL(msg string) {
	wrapper.Error(10006, msg)
}
