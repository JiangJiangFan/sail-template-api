package routers

import (
	"sail-chat/controllers"
	"sail-chat/global"
	m "sail-chat/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(m.CrossSite())
	r.GET("/test", controllers.TestDetail)
	r.GET("/testOne", controllers.TestSelectOne)
	r.PUT("/test", controllers.TestCreateUser)
	r.DELETE("/test", controllers.TestDelUser)
	// r.POST("/test", controllers.TestUpdateUser)
	r.POST("/login", controllers.Login)
	// r.GET("/info", controllers.GetInfo)
	r.GET("/user/all", m.JwtApis, controllers.TestDetail)
	r.POST("/user/register", controllers.Register)
	r.PUT("/user", m.JwtApis, m.TransMiddleware(global.App.DB), controllers.EditUser)
	r.GET("/city/all", controllers.GetCountryList)
	r.GET("/city/allByValue", controllers.GetCountryListByValue)

	return r
}
