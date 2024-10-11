package service

import (
	"sail-chat/models"
	"sail-chat/res"
	"sail-chat/types"
	"sail-chat/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAdmin(c *gin.Context, a types.Admin) {
	if tx, ok := c.Get("tx"); ok {
		db := tx.(*gorm.DB)
		hashedPass := utils.Sha256(a.Pass)
		a.Pass = hashedPass
		// 新建
		adminId, err := models.CreateAdmin(db, &a)
		if err != nil {
			res.Http(c).ErrorTransFail(err.Error())
			return
		}
		// 新建
		sectId, err := models.CreateSect(db, a.SectName)
		if err != nil {
			res.Http(c).ErrorTransFail(err.Error())
			return
		}
		// 关联
		if err := models.CreateAdminSect(db,adminId,sectId); err != nil {
			res.Http(c).ErrorTransFail(err.Error())
			return
		}
		res.Http(c).SuccessOnly(adminId)
	}
}