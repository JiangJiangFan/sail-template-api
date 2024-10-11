package controllers

import (
	"sail-chat/models"
	"sail-chat/res"
	"sail-chat/service"
	"sail-chat/types"

	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(c *gin.Context) {
	u := types.User{}
	err := c.ShouldBind(&u)
	if err != nil {
		res.Http(c).ErrorParam()
	}
	user, _ := models.GetUserByName(u.Username)
	if user.Username == "" {
		service.AddUser(c, u)
	} else {
		res.Http(c).ErrorLoginExist()
	}
}

// EditUser 修改用户信息
func EditUser(c *gin.Context){
	u := types.User{}
	if err := c.ShouldBind(&u); err != nil {
		res.Http(c).ErrorParam()
		return
	}
	if u.NewPass != "" {
		service.UpdatePass(c, u)
	} else {
		service.UpdateUser(c, u)
	}
}
