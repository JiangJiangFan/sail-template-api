package controllers

import (
	"sail-chat/res"
	"sail-chat/types"

	"github.com/gin-gonic/gin"
)

func GetCountryList(c *gin.Context) {
	// 获取国家列表的逻辑
	var meta *types.Meta
	err := c.ShouldBind(&meta)
	if err != nil {
		res.Http(c).ErrorContentType()
	} else {
		// list, e := models.SelectCountryCustom(meta)
		// if e != nil {
			res.Http(c).ErrorQuery()
		// } else {
		// 	res.Http(c).Success(list, meta)
		// }
	}
}

func GetCountryListByValue(c *gin.Context) {
	var meta *types.Meta
	err := c.ShouldBind(&meta)
	if err != nil {
		res.Http(c).ErrorContentType()
	} else {
		// list, e := models.SelectCountryCustomByValue(meta)
		// if e != nil {
			res.Http(c).ErrorQuery()
		// } else {
		// 	res.Http(c).Success(list, meta)
		// }
	}
}
