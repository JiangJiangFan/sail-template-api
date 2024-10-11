package middleware

import (
	"sail-chat/res"
	tool "sail-chat/utils"
	"time"

	"github.com/gin-gonic/gin"
)


func JwtApis(c *gin.Context) {
	// 获取token
	token := c.GetHeader("Authorization")

	if token == "" {
		token = c.Query("token")
	}
	// 解析token
	claims := tool.ValidateJwtToken(token)

	if claims == nil || claims["username"] == nil || claims["create_time"] == nil {
		res.Http(c).ErrorParse()
		return
	}
	createTime := int64(claims["create_time"].(float64))
	var expire int64 = 24 * 60 * 60
	now := time.Now().Unix()
	if (now - createTime) >= expire {
		res.Http(c).ErrorLoginToken()
		return
	}
}