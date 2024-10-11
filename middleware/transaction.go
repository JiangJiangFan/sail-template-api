package middleware

import (
	"fmt"
	"sail-chat/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func (c *gin.Context) {
		// 创建事务
		tx := db.Begin()
		// 存储事务对象到上下文
		c.Set("tx",tx)
		c.Next()
		if value, exists := c.Get("error"); exists {
			// 请求错误，回滚
			if value != nil {
				if err := tx.Rollback().Error; err != nil {
					res.Http(c).ErrorTrans(err.Error())
					return
				}
			}
			fmt.Println("事务：",value)
		} else {
			tx.Commit()
		}
	}
}
