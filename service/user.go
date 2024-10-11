package service

import (
	"sail-chat/models"
	"sail-chat/res"
	"sail-chat/types"
	"sail-chat/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddUser(c *gin.Context, u types.User) {
	if tx, ok := c.Get("tx"); ok {
		db := tx.(*gorm.DB)
		// 加密
		hashedPass := utils.Sha256(u.Password)
		u.Password = hashedPass
		// 新建用户
		id, err := models.CreateUser(db, &u)
		if err != nil {
			res.Http(c).ErrorTransFail(err.Error())
			return
		}
		// 设置权限
		if err := models.CreateUserRole(db, id, u.RoleId); err != nil {
			res.Http(c).ErrorTransFail(err.Error())
			return
		}
		res.Http(c).SuccessOnly(id)
	}
}
// 修改密码
func UpdatePass(c *gin.Context, u types.User) {
	user, _ := models.GetUserByName(u.Username)
	// 如果用户名为空
	if user.Username == "" {
		res.Http(c).ErrorLoginNot()
		return 
	}
	// 如果密码 不等于 数据库密码
		if user.Password != utils.Sha256(u.Password) {
			res.Http(c).ErrorLoginPass()
			return 
		}
		user.Password = utils.Sha256(u.NewPass)
	// 修改密码
	_, err := models.UpdatePass(user)
	res.Http(c).ErrorSQL(err.Error())
}

func UpdateUser(c *gin.Context, u types.User) {
	user, err := models.GetUserByID(u.Id)
	if err != nil {
		res.Http(c).ErrorSQL(err.Error())
		return
	}
	user.Username = u.Username
	user.Avator = u.Avator
	user.Nickname = u.Nickname
	if tx, ok := c.MustGet("tx").(*gorm.DB); ok {
		// 更新用户表
		if err := models.UpdateUser(tx, &user); err != nil {
			res.Http(c).ErrorSQL(err.Error())
			return
		}
		if _,err := models.GetRole(tx, u.RoleId); err != nil {
			res.Http(c).ErrorSQL(err.Error())
			return
		}
		// 更新中间表
		if err := models.UpdateUserRoleByUserId(tx,u.Id, u.RoleId); err != nil {
			res.Http(c).ErrorSQL(err.Error())
		} else {
			res.Http(c).SuccessOnly(nil)
		}
	} else {
		res.Http(c).ErrorParse()
	}
}