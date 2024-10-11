package controllers

import (
	"sail-chat/models"
	"sail-chat/res"
	"sail-chat/types"
	tool "sail-chat/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// 获取请求参数
	u := types.User{}
	// 验证参数
	if err := c.ShouldBind(&u); err != nil {
		msg := tool.ValidatorDefault(&u)
		res.Http(c).ErrorValidator(msg)
		return
	}
	user, _ := models.GetUserByName(u.Username)
	// 如果用户存在，则验证密码是否正确
	if user.Username != "" {
		// 加密
		hashedPass := tool.Sha256(u.Password)
		if user.Password == hashedPass {
			claims := make(map[string]interface{})
			claims["username"] = user.Username
			claims["nickname"] = user.Nickname
			claims["avator"] = user.Avator
			claims["create_time"] = time.Now().Unix()
			// claims["r_name"] = user.RoleName
			// claims["r_id"] = user.RoleId
			claims["u_id"] = user.Id
			// 生成token
			token, _ := tool.JwtToken(claims)
			res.Http(c).SuccessOnly(token)
		} else {
			res.Http(c).ErrorLoginPass()
		}
	} else {
		res.Http(c).ErrorLoginNot()
	}
}

// func GetInfo(c *gin.Context) {
// 	// 获取token
// 	token := c.GetHeader("Authorization")
// 	// 解析token
// 	claims, err := tool.ValidateJwtToken(token)
// 	if err != nil {
// 		res.Http(c).ErrorLoginToken()
// 		return
// 	}
// 	// 获取用户信息
// 	user, _ := models.GetUserByName(claims)
// 	res.Http(c).SuccessOnly(user)
// }
