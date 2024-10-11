package controllers

import (
	"sail-chat/models"
	res "sail-chat/res"
	"sail-chat/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestSelectOne(c *gin.Context) {
	var meta *types.Meta
	err := c.ShouldBind(&meta)
	if err != nil {
		res.Http(c).ErrorContentType()
	} else {
		user, _ := models.GetUserByCustom(meta.Value)
		res.Http(c).Success(user, meta)
	}
}

// TestDetail 查询
func TestDetail(c *gin.Context) {
	meta := types.Meta{}
	err := c.ShouldBind(&meta)
	if err != nil {
		res.Http(c).ErrorContentType()
		return
	} else {
		user, count, e := models.GetUsers(&meta)
		if e != nil {
			res.Http(c).ErrorQuery()
			return
		}
		meta.Total = count
		// // 测试详情页
		res.Http(c).Success(&user, &meta)
	}
}

// TestCreateUser 创建
func TestCreateUser(c *gin.Context) {
	tx, ok := c.Get("tx")
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	user := models.GetUserByNameAndPassword(name, pass)
	u := &types.User{
		Username: name,
		Password: pass,
	}
	if user.Username == "" && ok {
		db := tx.(*gorm.DB)
		id,_ := models.CreateUser(db, u)
		res.Http(c).SuccessOnly(id)
	} else {
		res.Http(c).ErrorLoginExist()
	}
}

// TestDelUser 删除
func TestDelUser(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	user := models.GetUserByNameAndPassword(name, pass)
	if user.Username == "" {
		res.Http(c).ErrorLoginNot()
	} else {
		err := models.DeleteUser(user)
		if err != nil {
			return
		}
		res.Http(c).SuccessOnly("删除成功")
	}
}

// TestUpdateUser 修改
// func TestUpdateUser(c *gin.Context) {
// 	name := c.PostForm("name")
// 	pass := c.PostForm("pass")
// 	user := models.GetUserByNameAndPassword(name, pass)
// 	if user.Username == "" {
// 		res.Http(c).ErrorLoginNot()
// 	} else {
// 		u := &types.User{
// 		Id: user.Id,
// 		Username: name,
// 		Password: pass,
// 	}
// 		err := models.UpdateUser(u)
// 		if err != nil {
// 			return
// 		}
// 		res.Http(c).SuccessOnly("修改成功")
// 	}
// }
